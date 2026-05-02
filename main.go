package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

const catAscii = `
    /\_____/\
   /  o   o  \
  ( ==  ^  == )
   )         (
  (           )
 ( (  )   (  ) )
(__(__)___(__)__)
`

func wrapLines(lines []string, maxWidth int) []string {
	var result []string
	for _, line := range lines {
		runes := []rune(line)
		for len(runes) > maxWidth {
			result = append(result, string(runes[:maxWidth]))
			runes = runes[maxWidth:]
		}
		result = append(result, string(runes))
	}
	return result
}

func buildBubble(lines []string) string {
	maxLen := 0
	for _, l := range lines {
		if w := utf8.RuneCountInString(l); w > maxLen {
			maxLen = w
		}
	}

	var sb strings.Builder
	border := strings.Repeat("-", maxLen+2)
	sb.WriteString(" " + border + "\n")
	for i, l := range lines {
		pad := maxLen - utf8.RuneCountInString(l)
		left, right := "| ", " |"
		if len(lines) == 1 {
			left, right = "< ", " >"
		} else if i == 0 {
			left, right = "/ ", " \\"
		} else if i == len(lines)-1 {
			left, right = "\\ ", " /"
		}
		sb.WriteString(left + l + strings.Repeat(" ", pad) + right + "\n")
	}
	sb.WriteString(" " + border + "\n")
	return sb.String()
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: catsay [-s|--string \"text\"] [file ...]")
	fmt.Fprintln(os.Stderr, "       echo text | catsay")
	os.Exit(1)
}

func main() {
	var lines []string
	wrap := 60

	args := os.Args[1:]

	// parse flags
	var fileArgs []string
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-s", "--string":
			if i+1 >= len(args) {
				fmt.Fprintln(os.Stderr, "catsay: -s requires an argument")
				usage()
			}
			i++
			lines = append(lines, strings.Split(args[i], "\n")...)
		case "-h", "--help":
			usage()
		default:
			fileArgs = append(fileArgs, args[i])
		}
	}

	// read files if any were given
	for _, arg := range fileArgs {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "catsay: %s: %v\n", arg, err)
			continue
		}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		f.Close()
	}

	// no args at all: print usage (don't block on stdin)
	if len(os.Args) == 1 {
		usage()
	}

	// no content from flags/files but stdin was piped
	if len(lines) == 0 && len(fileArgs) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	}

	if len(lines) == 0 {
		lines = []string{"..."}
	}

	lines = wrapLines(lines, wrap)
	fmt.Print(buildBubble(lines))
	fmt.Print(catAscii)
}
