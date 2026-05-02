#!/bin/sh
set -e

BIN=catsay
REPO="LiTLiTschi/catsay"
INSTALL_DIR="/usr/local/bin"

# detect arch
ARCH=$(uname -m)
OS=$(uname -s | tr '[:upper:]' '[:lower:]')

case "$ARCH" in
  x86_64)  ARCH="amd64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

case "$OS" in
  linux|darwin) ;;
  *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

SUFFIX="${OS}-${ARCH}"

echo "Fetching latest release for ${OS}/${ARCH}..."

# get latest tag
LATEST=$(curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest" \
  | grep '"tag_name"' | head -1 | sed 's/.*"tag_name": *"\([^"]*\)".*/\1/')

if [ -z "$LATEST" ]; then
  echo "Could not determine latest release. Try: go install github.com/${REPO}@latest"
  exit 1
fi

URL="https://github.com/${REPO}/releases/download/${LATEST}/${BIN}-${SUFFIX}"

echo "Downloading ${BIN} ${LATEST}..."
curl -fsSL "$URL" -o "/tmp/${BIN}"
chmod +x "/tmp/${BIN}"

# install (try sudo, fall back to ~/.local/bin)
if [ -w "$INSTALL_DIR" ]; then
  mv "/tmp/${BIN}" "${INSTALL_DIR}/${BIN}"
elif command -v sudo >/dev/null 2>&1; then
  sudo mv "/tmp/${BIN}" "${INSTALL_DIR}/${BIN}"
else
  mkdir -p "$HOME/.local/bin"
  mv "/tmp/${BIN}" "$HOME/.local/bin/${BIN}"
  INSTALL_DIR="$HOME/.local/bin"
fi

echo "Installed to ${INSTALL_DIR}/${BIN}"
${INSTALL_DIR}/${BIN} --version 2>/dev/null || true
echo "Done! Run: catsay <file>"
