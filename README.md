# 🐱 catsay

> Like `cowsay`, but it's a **cat** — and it reads files directly, like `cat`.

```
 ----------------------------------------------------
/ package main                                      \
| import "fmt"                                       |
\ func main() { fmt.Println("meow") }               /
 ----------------------------------------------------
    /\_____/\
   /  o   o  \
  ( ==  ^  == )
   )         (
  (           )
 ( (  )   (  ) )
(__(__)___(__)__)
```

## Install

### Linux / macOS — curl

```sh
curl -fsSL https://raw.githubusercontent.com/LiTLiTschi/catsay/main/install.sh | sh
```

### Linux / macOS — wget

```sh
wget -qO- https://raw.githubusercontent.com/LiTLiTschi/catsay/main/install.sh | sh
```

No repo clone needed. The script:
- Works with either `curl` or `wget` (whichever is available)
- Downloads a fully static prebuilt binary for your arch (`amd64` / `arm64`)
- Falls back to building from source if no release exists yet (requires Go)
- Installs to `/usr/local/bin/` — or `~/.local/bin/` if you don’t have sudo

### Windows — PowerShell

```powershell
irm https://raw.githubusercontent.com/LiTLiTschi/catsay/main/install.ps1 | iex
```

Installs to `%LOCALAPPDATA%\Programs\catsay\` and adds it to your user PATH automatically.

### go install (if you have Go)

```bash
go install github.com/LiTLiTschi/catsay@latest
```

## Usage

```bash
# Read a file directly (the point of this tool)
catsay README.md

# Multiple files — concatenated like real cat
catsay file1.txt file2.txt

# Stdin fallback (pipe still works)
echo "meow" | catsay
cat somefile | catsay

# Empty? Cat says ...
catsay
```

## Why?

`cowsay` needs its text piped in. `cat` reads files directly. `catsay` fuses both:
you give it a file (or a few), it reads them and has the cat say the contents.

The binary is fully static — no libc, no runtime, no nothing. Drop it anywhere and it runs.

## How it differs from cowsay

| | `cowsay` | `catsay` |
|---|---|---|
| Animal | 🐄 Cow | 🐱 Cat |
| Input | stdin only | **file args** + stdin fallback |
| Multiple files | ❌ | ✅ concatenated |
| Dependencies | Perl | **none** |
| Install | package manager | one `curl` or `wget` command |

## License

MIT
