package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

// expandTabs replaces tab characters with spaces (tabstop=8) so that
// rune counting reflects the visual column width.
func expandTabs(s string) string {
	var b strings.Builder
	col := 0
	for _, r := range s {
		if r == '\t' {
			spaces := 8 - (col % 8)
			b.WriteString(strings.Repeat(" ", spaces))
			col += spaces
		} else {
			b.WriteRune(r)
			col++
		}
	}
	return b.String()
}

// buildCat generates the ASCII cat widened by fatness factor.
func buildCat(fat int) string {
	if fat < 1 {
		fat = 1
	}
	g := (fat - 1) * 2

	headU := strings.Repeat("_", 5+g)
	eyeGap := strings.Repeat(" ", 3+g)
	noseHalf := strings.Repeat(" ", 2+g/2)
	body1 := strings.Repeat(" ", 9+g)
	body2 := strings.Repeat(" ", 11+g)
	pawGap := strings.Repeat(" ", 3+g)
	feetU := strings.Repeat("_", 3+g)

	return fmt.Sprintf(
		"\n"+
			"    /\\%s/\\\n"+
			"   /  o%so  \\\n"+
			"  ( ==%s^%s== )\n"+
			"   )%s(\n"+
			"  (%s)\n"+
			" ( (  )%s(  ) )\n"+
			"(__(__)%s(__)__)\n",
		headU,
		eyeGap,
		noseHalf, noseHalf,
		body1,
		body2,
		pawGap,
		feetU,
	)
}

// wrapLines hard-wraps lines to maxWidth runes.
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
	fmt.Fprintln(os.Stderr, "Usage: catsay [-s|--string \"text\"] [-f|--fat N] [file ...]")
	fmt.Fprintln(os.Stderr, "       echo text | catsay")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "  -s, --string  say a string directly")
	fmt.Fprintln(os.Stderr, "  -f, --fat N   scale cat width (default: 1)")
	os.Exit(1)
}

func main() {
	var lines []string
	fat := 1

	args := os.Args[1:]

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
		case "-f", "--fat":
			if i+1 >= len(args) {
				fmt.Fprintln(os.Stderr, "catsay: -f requires a number")
				usage()
			}
			i++
			n, err := strconv.Atoi(args[i])
			if err != nil || n < 1 {
				fmt.Fprintln(os.Stderr, "catsay: -f requires a positive integer")
				usage()
			}
			fat = n
		case "-h", "--help":
			usage()
		default:
			fileArgs = append(fileArgs, args[i])
		}
	}

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

	if len(os.Args) == 1 {
		usage()
	}

	if len(lines) == 0 && len(fileArgs) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	}

	if len(lines) == 0 {
		lines = []string{"..."}
	}

	for i, l := range lines {
		lines[i] = expandTabs(l)
	}

	wrap := termWidth() - 4
	// Cap the wrap width to prevent excessively wide borders in some terminal environments
	if wrap > 120 {
		wrap = 120
	}
	if wrap < 20 {
		wrap = 20
	}

	lines = wrapLines(lines, wrap)
	fmt.Print(buildBubble(lines))
	fmt.Print(buildCat(fat))
}
