package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/huytd/stackoverlazy/colors"
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
	searchResponse := search.Query(query)
	if searchResponse != nil {

	} else {
		fmt.Println(colors.Apply("<red>Unable to search for answer! Please try again!</red>"))
	}
}
