package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

const catAscii = `
    \
     \
      /\_____/\
     ( o   o  )
     =( Y  )=
      )     (
     (_)-(_)
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

func main() {
	var lines []string
	wrap := 60

	args := os.Args[1:]
	if len(args) == 0 {
		// read from stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	} else {
		for _, arg := range args {
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
	}

	if len(lines) == 0 {
		lines = []string{"..."}
	}

	lines = wrapLines(lines, wrap)
	fmt.Print(buildBubble(lines))
	fmt.Print(catAscii)
}
