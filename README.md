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

### One-liner (Linux & macOS)

```sh
curl -fsSL https://raw.githubusercontent.com/LiTLiTschi/catsay/main/install.sh | sh
```

No dependencies required — just `curl`. Downloads a fully static binary for your arch (`amd64` or `arm64`).
Installs to `/usr/local/bin/catsay`, or `~/.local/bin/catsay` if you don't have sudo.

### go install (if you have Go)

```bash
go install github.com/LiTLiTschi/catsay@latest
```

### From source

```bash
git clone https://github.com/LiTLiTschi/catsay
cd catsay
CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o catsay .
sudo mv catsay /usr/local/bin/
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
| Install | package manager | `curl \| sh` or single binary |

## License

MIT
