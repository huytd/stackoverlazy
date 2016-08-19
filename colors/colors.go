package colors

import (
	"fmt"
	"regexp"
)

var (
	TERM_STYLE = []string{
		// LIGHT
		"\033[1m",
		// UNDERLINE
		"\033[4m",
		// YELLOW
		"\033[33m",
		// RED
		"\033[31m",
		// GREEN
		"\033[32m",
		// CYAN
		"\033[36m",
		// BLUE
		"\033[34m",
	}

	TERM_RESET = "\033[0m"

	INPUT_TAGS = []string{
		"light",
		"u",
		"yellow",
		"red",
		"green",
		"cyan",
		"blue",
	}
)

func Apply(input string) string {
	out := input
	for i := 0; i < len(INPUT_TAGS); i++ {
		re := regexp.MustCompile(fmt.Sprintf(`(?s)\<%s\>(.*?)<\/%s\>`, INPUT_TAGS[i], INPUT_TAGS[i]))
		out = re.ReplaceAllString(out, TERM_STYLE[i]+"$1"+TERM_RESET)
	}
	return out
}
