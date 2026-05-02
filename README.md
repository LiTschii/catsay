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

## Why?

`cowsay` needs its text piped in. `cat` reads files directly. `catsay` fuses both:
you give it a file (or a few), it reads them and has the cat say the contents.

## Install

```bash
# From source (requires Go 1.21+)
git clone https://github.com/LiTLiTschi/catsay
cd catsay
go build -o catsay .
sudo mv catsay /usr/local/bin/
```

Or with `go install`:

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

## How it differs from cowsay

| | `cowsay` | `catsay` |
|---|---|---|
| Animal | 🐄 Cow | 🐱 Cat |
| Input | stdin only | **file args** + stdin fallback |
| Multiple files | ❌ | ✅ concatenated |
| Install | needs package manager | single Go binary |

## Long lines

Lines longer than 60 characters are automatically wrapped. Multi-line content
gets the classic cowsay-style `/ ... \` border corners; single-line gets `< ... >`.

## License

MIT
