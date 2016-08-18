package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/huytd/stackoverlazy/colors"
	"github.com/huytd/stackoverlazy/parser"
	"github.com/huytd/stackoverlazy/search"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(colors.Apply("<red>No search term provided.</red>\n<red>Example:</red>\n\n  <yellow>stackoverlazy</yellow> <green>css vertical align</green>\n\n"))
		return
	}
	query := "stackoverflow+"
	query += strings.Join(args[:], "+")
	fmt.Print("Looking for the best answer")
	searchResponse := search.Query("https://search.yahoo.com/search?p=" + query)
	if searchResponse != nil {
		stackOverflowURL := parser.ParseURL(searchResponse)
		if stackOverflowURL != "" {
			stackOverflowResponse := search.Query(stackOverflowURL)
			if stackOverflowResponse != nil {
				answer := parser.ParseAnswer(stackOverflowResponse)
				fmt.Println("\n")
				fmt.Println(colors.Apply(answer))
				return
			}
		}
		fmt.Println("\n")
		fmt.Println(colors.Apply("<red>No answer found! Sorry buddy! You have to solve it yourself...</red>"))
		return
	}
	fmt.Println("\n")
	fmt.Println(colors.Apply("<red>Unable to search for answer! Please try again!</red>"))
}
