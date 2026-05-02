# catsay installer for Windows (PowerShell)
# Usage: irm https://raw.githubusercontent.com/LiTLiTschi/catsay/main/install.ps1 | iex

$ErrorActionPreference = 'Stop'

$Repo   = 'LiTLiTschi/catsay'
$Bin    = 'catsay.exe'

# detect arch
$Arch = if ([System.Environment]::Is64BitOperatingSystem) {
  if ($env:PROCESSOR_ARCHITECTURE -eq 'ARM64') { 'arm64' } else { 'amd64' }
} else {
  Write-Error 'Only 64-bit Windows is supported.'; exit 1
}

$Suffix = "windows-$Arch.exe"

# resolve install dir — prefer a dir already on PATH, else create one
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

$Url  = "https://github.com/$Repo/releases/download/$Tag/catsay-$Suffix"
$Dest = Join-Path $InstallDir $Bin

Write-Host "Downloading catsay $Tag ($Arch)..."
Invoke-WebRequest -Uri $Url -OutFile $Dest -UseBasicParsing

Write-Host "Installed -> $Dest"
Write-Host "Run: catsay <file>"
Write-Host "(You may need to restart your terminal for PATH to take effect.)"
