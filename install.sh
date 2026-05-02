#!/bin/sh
set -e

BIN=catsay
REPO="LiTLiTschi/catsay"
INSTALL_DIR="/usr/local/bin"
TMP=$(mktemp)

die() { echo "error: $1" >&2; exit 1; }

# pick downloader
if command -v curl >/dev/null 2>&1; then
  fetch() { curl -fsSL "$1" -o "$2"; }
elif command -v wget >/dev/null 2>&1; then
  fetch() { wget -qO "$2" "$1"; }
else
  die "neither curl nor wget found"
fi

# detect OS / arch
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
case "$ARCH" in
  x86_64)        ARCH=amd64 ;;
  aarch64|arm64) ARCH=arm64 ;;
  *) die "unsupported architecture: $ARCH" ;;
esac
case "$OS" in
  linux|darwin) ;;
  *) die "unsupported OS: $OS" ;;
esac

SUFFIX="${OS}-${ARCH}"

# resolve install dir
if [ ! -w "$INSTALL_DIR" ]; then
  if command -v sudo >/dev/null 2>&1; then
    SUDO=sudo
  else
    INSTALL_DIR="$HOME/.local/bin"
    mkdir -p "$INSTALL_DIR"
  fi
fi

# --- try prebuilt binary from latest release ---
LATEST=$(fetch "https://api.github.com/repos/${REPO}/releases/latest" /dev/stdout 2>/dev/null \
  | grep '"tag_name"' | head -1 \
  | sed 's/.*"tag_name": *"\([^"]*\)".*/\1/')

if [ -n "$LATEST" ]; then
  URL="https://github.com/${REPO}/releases/download/${LATEST}/${BIN}-${SUFFIX}"
  echo "Downloading ${BIN} ${LATEST} (${SUFFIX})..."
  if fetch "$URL" "$TMP" 2>/dev/null && [ -s "$TMP" ]; then
    chmod +x "$TMP"
    $SUDO mv "$TMP" "${INSTALL_DIR}/${BIN}"
    echo "Installed ${BIN} ${LATEST} -> ${INSTALL_DIR}/${BIN}"
    exit 0
  fi
fi

# --- fallback: build from source with Go ---
echo "No prebuilt binary found. Trying to build from source..."
command -v go >/dev/null 2>&1 || die "Go is not installed and no prebuilt binary is available.\nInstall Go from https://go.dev/dl/ and re-run, or download a binary from:\nhttps://github.com/${REPO}/releases"

BUILDTMP=$(mktemp -d)
trap 'rm -rf "$BUILDTMP"' EXIT

fetch "https://raw.githubusercontent.com/${REPO}/main/main.go" "${BUILDTMP}/main.go"
fetch "https://raw.githubusercontent.com/${REPO}/main/go.mod"  "${BUILDTMP}/go.mod"

( cd "$BUILDTMP" && CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o "$TMP" . )
chmod +x "$TMP"
$SUDO mv "$TMP" "${INSTALL_DIR}/${BIN}"
echo "Built and installed ${BIN} -> ${INSTALL_DIR}/${BIN}"
