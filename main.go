package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/huytd/stackoverlazy/colors"
	"github.com/huytd/stackoverlazy/parser"
	"github.com/huytd/stackoverlazy/search"
	"github.com/mattn/go-colorable"
)

func main() {
	out := colorable.NewColorableStdout()

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintln(out, colors.Apply("<red>No search term provided.</red>\n<red>Example:</red>\n\n  <yellow>stackoverlazy</yellow> <green>css vertical align</green>\n\n"))
		return
	}
	query := "stackoverflow+"
	query += strings.Join(args[:], "+")
	fmt.Fprint(out, "Looking for the best answer")
	searchResponse := search.Query("https://google.com/search?q=" + query)
	if searchResponse != nil {
		stackOverflowURL := parser.ParseURL(searchResponse)
		if stackOverflowURL != "" {
			stackOverflowResponse := search.Query(stackOverflowURL)
			if stackOverflowResponse != nil {
				answer := parser.ParseAnswer(stackOverflowResponse)
				fmt.Fprintln(out, "\n")
				fmt.Fprintln(out, colors.Apply(answer))
				fmt.Fprintln(out, colors.Apply("\n<green><u>See more:</u></green> <blue>"+stackOverflowURL+"</blue>"))
				return
			}
		}
		fmt.Fprintln(out, "\n")
		fmt.Fprintln(out, colors.Apply("<red>No answer found! Sorry buddy! You have to solve it yourself...</red>"))
		return
	}
	fmt.Fprintln(out, "\n")
	fmt.Fprintln(out, colors.Apply("<red>Unable to search for answer! Please try again!</red>"))
}
