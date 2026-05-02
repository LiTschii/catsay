#!/bin/sh
set -e

BIN=catsay
REPO="LiTLiTschi/catsay"
INSTALL_DIR="/usr/local/bin"
TMP=$(mktemp)
TMPTAR=$(mktemp)

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
  URL="https://github.com/${REPO}/releases/download/${LATEST}/${BIN}-${SUFFIX}.tar.gz"
  echo "Downloading ${BIN} ${LATEST} (${SUFFIX})..."
  if fetch "$URL" "$TMPTAR" 2>/dev/null && [ -s "$TMPTAR" ]; then
    tar -xzf "$TMPTAR" -C "$(dirname "$TMP")" "$BIN" 2>/dev/null \
      || tar -xzf "$TMPTAR" -C "$(dirname "$TMP")" --strip-components=1 2>/dev/null
    EXTRACTED="$(dirname "$TMP")/$BIN"
    if [ -f "$EXTRACTED" ]; then
      chmod +x "$EXTRACTED"
      $SUDO mv "$EXTRACTED" "${INSTALL_DIR}/${BIN}"
    else
      # binary might be directly in the tarball root with a different layout
      tar -xzf "$TMPTAR" -O > "$TMP" 2>/dev/null
      chmod +x "$TMP"
      $SUDO mv "$TMP" "${INSTALL_DIR}/${BIN}"
    fi
    rm -f "$TMPTAR"
    echo "Installed ${BIN} ${LATEST} -> ${INSTALL_DIR}/${BIN}"
    exit 0
  fi
  rm -f "$TMPTAR"
fi

# --- fallback: build from source with Go ---
echo "No prebuilt binary found. Falling back to building from source..."

if ! command -v go >/dev/null 2>&1; then
  echo "Go is not installed."
  printf "Install Go now? (Y/n) "
  read -r REPLY </dev/tty
  case "$REPLY" in
    [nN]*) die "Aborting. Install Go from https://go.dev/dl/ or wait for a prebuilt release at https://github.com/${REPO}/releases" ;;
  esac

  echo "Installing Go via the official installer..."
  GOTMP=$(mktemp -d)
  trap 'rm -rf "$GOTMP"' EXIT
  GO_VERSION="1.22.3"
  GO_TARBALL="go${GO_VERSION}.${OS}-${ARCH}.tar.gz"
  fetch "https://go.dev/dl/${GO_TARBALL}" "${GOTMP}/${GO_TARBALL}"
  $SUDO tar -C /usr/local -xzf "${GOTMP}/${GO_TARBALL}"
  export PATH="$PATH:/usr/local/go/bin"
  echo "Go installed to /usr/local/go"
fi

BUILDTMP=$(mktemp -d)
trap 'rm -rf "$BUILDTMP"' EXIT

fetch "https://raw.githubusercontent.com/${REPO}/main/main.go" "${BUILDTMP}/main.go"
fetch "https://raw.githubusercontent.com/${REPO}/main/go.mod"  "${BUILDTMP}/go.mod"

( cd "$BUILDTMP" && CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o "$TMP" . )
chmod +x "$TMP"
$SUDO mv "$TMP" "${INSTALL_DIR}/${BIN}"
echo "Built and installed ${BIN} -> ${INSTALL_DIR}/${BIN}"
