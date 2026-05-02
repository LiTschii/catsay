# catsay installer for Windows (PowerShell)
# Usage: irm https://raw.githubusercontent.com/LiTLiTschi/catsay/main/install.ps1 | iex

$ErrorActionPreference = 'Stop'

$Repo = 'LiTLiTschi/catsay'
$Bin  = 'catsay.exe'

# resolve install dir
$InstallDir = "$env:LOCALAPPDATA\Programs\catsay"
if (-not (Test-Path $InstallDir)) { New-Item -ItemType Directory -Path $InstallDir | Out-Null }

# add to user PATH if not already there
$UserPath = [System.Environment]::GetEnvironmentVariable('PATH', 'User')
if ($UserPath -notlike "*$InstallDir*") {
  [System.Environment]::SetEnvironmentVariable('PATH', "$UserPath;$InstallDir", 'User')
  $env:PATH += ";$InstallDir"
  Write-Host "Added $InstallDir to PATH."
}

# get latest release tag
$Release = Invoke-RestMethod "https://api.github.com/repos/$Repo/releases/latest"
$Tag     = $Release.tag_name

if (-not $Tag) {
  Write-Error "Could not find a release. Try: go install github.com/$Repo@latest"
  exit 1
}

# construct URL directly from tag — no asset list lookup needed
$Url  = "https://github.com/$Repo/releases/download/$Tag/$Bin"
$Dest = Join-Path $InstallDir $Bin

Write-Host "Downloading catsay $Tag..."
try {
  Invoke-WebRequest -Uri $Url -OutFile $Dest -UseBasicParsing
  Write-Host "Installed -> $Dest"
} catch {
  Write-Host "No prebuilt binary found for $Tag. Falling back to go install..."
  if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    Write-Error "Go is not installed. Install it from https://go.dev/dl/ and re-run this script, or download catsay manually from https://github.com/$Repo/releases"
    exit 1
  }
  $env:GOBIN = $InstallDir
  go install "github.com/$Repo@latest"
  Write-Host "Installed -> $InstallDir\$Bin"
}

Write-Host "Run: catsay <file>"
Write-Host "(You may need to restart your terminal for PATH to take effect.)"
