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
- Installs to `/usr/local/bin/` — or `~/.local/bin/` if you don't have sudo

### Windows — PowerShell

```powershell
irm https://raw.githubusercontent.com/LiTLiTschi/catsay/main/install.ps1 | iex
```

Installs to `%LOCALAPPDATA%\Programs\catsay\` and adds it to your user PATH automatically.

### go install (if you have Go)

```bash
go install github.com/LiTLiTschi/catsay@latest
```

### Optional but Recommended: replace `cat` with `catsay`

`catsay` is a strict superset of `cat` for reading files — it accepts the same file arguments and
falls back to stdin just like `cat` does, so aliasing it is safe in interactive shells.
The upside: every time you inspect a config file, a log, or a script, you get immediate visual
confirmation that the output belongs to a single command invocation, which is useful when
your terminal history is cluttered or you’re piping output from multiple sources.

Add to your `~/.bashrc` or `~/.zshrc`:

```sh
alias cat='catsay'
```

Then reload your shell:

```sh
source ~/.bashrc   # or source ~/.zshrc
```

> **Note:** The alias only applies to interactive shells. Scripts that call `cat` directly are
> unaffected, so nothing breaks.

## Usage

```bash
# Read a file directly
catsay README.md

# Multiple files — concatenated like real cat
catsay file1.txt file2.txt

# Pass a string directly
catsay -s "hello world"

# Stdin fallback (pipe still works)
echo "meow" | catsay

# Empty? Cat says ...
catsay
```

## Flags

### `-s`, `--string`

Passes a string directly as the cat's message, without needing a file or a pipe.
Useful when you want to quickly surface a value, a variable, or a short note in your
terminal output without creating a temporary file or a subshell.

```bash
catsay -s "build succeeded"
catsay -s "$MY_ENV_VAR"
```

### `-f`, `--fat N`

Scales the cat's body width by `N`. Default is `1`.

The horizontal scale factor directly controls the girth of the rendered cat.
This is particularly useful in wide terminal environments where the default
cat renders as visually undersized relative to the speech bubble — a mismatch
that can make the output harder to parse at a glance. A value of `2` or `3`
produces a more proportional result on terminals wider than 120 columns.

```bash
catsay -s "i am normal"          # default
catsay -s "i had too much food" -f 3
```

```
    /\_____/\                       /\_________/\
   /  o   o  \                     /  o     o  \
  ( ==  ^  == )     -f 3 ->       ( ==    ^    == )
   )         (                     )             (
  (           )                   (               )
 ( (  )   (  ) )                 ( (   )   (   ) )
(__(__)___(__)__)               (__(__)___(__)__)
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

## FAQ

**Q: Does the cat look like Garfield?**

No. The cat is a generic felid rendered in standard ASCII box-drawing conventions.
Garfield is a trademarked character with a distinct oval head, closed eyes, and
visible stripes. The `catsay` cat has a triangular ear profile, open circular eyes,
and a symmetric facial structure that is consistent with a generic domestic shorthair
representation. Any resemblance is a consequence of the limited character set available
in a 7-bit ASCII environment, not an intentional reference.

---

**Q: I used `-f 3` and now it looks even more like Garfield. Is this intentional?**

No. The `-f` flag increases horizontal body width to improve proportionality on wide
terminals, as documented above. The fact that a wider, rounder cat body superficially
resembles a well-known overweight cartoon cat is an unavoidable geometric consequence
of scaling a symmetric ASCII figure outward. This is not a design goal. It is a
collateral outcome.

---

**Q: My colleague said "hey it's Garfield" when they saw the output. Should I be concerned?**

No. This is a common reaction caused by pattern recognition bias — humans are highly
trained to identify cat-shaped forms and associate them with culturally prominent
cat characters. The `catsay` cat predates any specific pop-culture reference in the
sense that it is derived from first principles of ASCII art cat construction. Your
colleague's association, while understandable, is not legally or technically meaningful.

---

**Q: Could Paws, Inc. take legal action over the cat's appearance?**

This question falls outside the scope of a README. That said: ASCII art consisting
of standard punctuation characters arranged to suggest a round-faced cat does not
constitute a copyrightable or trademarkable likeness under any jurisdiction the
author is aware of. The character `(` is not owned by anyone.

---

**Q: Will there ever be a version of the cat that looks less like Garfield?**

The current cat design is considered stable. Proposals to alter the facial geometry
in order to distance it from any specific cultural reference are welcome via GitHub
issues, but should be accompanied by a concrete ASCII art alternative that still
reads unambiguously as a cat at default terminal font sizes. This is harder than it
sounds.

## License

MIT
