package parser_test

import (
	"testing"

	"github.com/huytd/stackoverlazy/parser"
)

func TestParseURLWithNilInput(t *testing.T) {
	expect := parser.ParseURL(nil)
	if expect != "" {
		t.Fail()
	}
}

func TestParseAnswerWithNilInput(t *testing.T) {
	expect := parser.ParseAnswer(nil)
	if expect != "" {
		t.Fail()
	}
}
