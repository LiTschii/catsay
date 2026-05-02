package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

// buildCat generates the ASCII cat scaled by fatness.
// fatness=1 is default. Each extra unit repeats the belly rows
// vertically and widens the body horizontally.
func buildCat(fat int) string {
	if fat < 1 {
		fat = 1
	}

	// horizontal padding added to body lines (spaces inside the sides)
	hpad := strings.Repeat(" ", (fat-1)*2)
	// extra belly rows repeated vertically
	belly := strings.Repeat("  ( "+hpad+"          "+hpad+" )\n", fat-1)

	return fmt.Sprintf(`
    /\_____/\
   /  o%s   o  \
  ( ==%s  ^  ==%s )
   )%s         %s(
%s  (  %s         %s  )
 ( ( %s )   ( %s ) )
(__(__)___(__)__)
`,
		hpad, hpad, hpad,
		hpad, hpad,
		belly,
		hpad, hpad,
		hpad, hpad,
	)
}

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
	fmt.Fprintln(os.Stderr, "  -f, --fat N   scale cat fatness (default: 1)")
	os.Exit(1)
}

func main() {
	var lines []string
	wrap := 60
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

	lines = wrapLines(lines, wrap)
	fmt.Print(buildBubble(lines))
	fmt.Print(buildCat(fat))
}
