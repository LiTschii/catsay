# catsay (PyPI)

This directory contains the Python packaging shim that makes `catsay` installable via `pip`.

The package downloads the correct prebuilt Go binary for your platform on first run and caches it in `~/.cache/catsay/<version>/`. All subsequent calls go directly to the binary — Python is not involved at runtime.

## Usage

```sh
pip install catsay
catsay README.md
```

## How it works

1. `pip install catsay` installs a thin Python shim and a `catsay` entry point script.
2. On first run, the shim detects your OS/arch, downloads the matching binary from the GitHub release, marks it executable, and caches it.
3. On every subsequent run, the shim calls `os.execv` (Unix) or `subprocess.run` (Windows) to hand off to the native binary directly.

The binary itself is a fully static Go binary — no libc, no runtime. The Python layer is purely a delivery mechanism.

## Publishing to PyPI

```sh
cd python/
pip install build twine
python -m build
twine upload dist/*
```

Make sure a GitHub release tagged `v<version>` exists with the prebuilt binaries before publishing.
