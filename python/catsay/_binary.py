"""Download and cache the catsay binary for the current platform."""

import os
import platform
import stat
import sys
import urllib.request
from pathlib import Path

VERSION = "0.1.0"
GITHUB_REPO = "LiTLiTschi/catsay"


def _platform_asset() -> str:
    """Return the GitHub release asset filename for the current platform."""
    system = platform.system().lower()
    machine = platform.machine().lower()

    arch = "arm64" if machine in ("arm64", "aarch64") else "amd64"

    if system == "linux":
        return f"catsay-linux-{arch}"
    elif system == "darwin":
        return f"catsay-darwin-{arch}"
    elif system == "windows":
        return f"catsay-windows-{arch}.exe"
    else:
        raise RuntimeError(
            f"Unsupported platform: {system}/{machine}. "
            "Please build from source: https://github.com/LiTLiTschi/catsay"
        )


def _cache_dir() -> Path:
    """Return (and create) the directory where the binary is cached."""
    base = Path(os.environ.get("XDG_CACHE_HOME", Path.home() / ".cache"))
    d = base / "catsay" / VERSION
    d.mkdir(parents=True, exist_ok=True)
    return d


def _binary_path() -> Path:
    asset = _platform_asset()
    ext = ".exe" if platform.system().lower() == "windows" else ""
    return _cache_dir() / f"catsay{ext}"


def ensure_binary() -> Path:
    """Return the path to the catsay binary, downloading it if necessary."""
    dest = _binary_path()
    if dest.exists():
        return dest

    asset = _platform_asset()
    url = f"https://github.com/{GITHUB_REPO}/releases/download/v{VERSION}/{asset}"

    print(f"catsay: downloading binary from {url}", file=sys.stderr)
    try:
        urllib.request.urlretrieve(url, dest)
    except Exception as exc:
        dest.unlink(missing_ok=True)
        raise RuntimeError(
            f"Failed to download catsay binary from {url}:\n  {exc}\n"
            "You can build from source instead: https://github.com/LiTLiTschi/catsay"
        ) from exc

    # Make executable on Unix
    if platform.system().lower() != "windows":
        dest.chmod(dest.stat().st_mode | stat.S_IEXEC | stat.S_IXGRP | stat.S_IXOTH)

    return dest
