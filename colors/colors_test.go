package colors_test

import (
	"testing"

	"github.com/huytd/stackoverlazy/colors"
)

func TestColorApplyWithRed(t *testing.T) {
	input := "<red>hello</red>"
	expect := "\033[31mhello\033[0m"
	if colors.Apply(input) != expect {
		t.Fail()
	}
}

func TestColorApplyWithYellow(t *testing.T) {
	input := "<yellow>hello</yellow>"
	expect := "\033[33mhello\033[0m"
	if colors.Apply(input) != expect {
		t.Fail()
	}
}

func TestColorApplyWithGreen(t *testing.T) {
	input := "<green>hello</green>"
	expect := "\033[32mhello\033[0m"
	if colors.Apply(input) != expect {
		t.Fail()
	}
}

func TestColorApplyWithCyan(t *testing.T) {
	input := "<cyan>hello</cyan>"
	expect := "\033[36mhello\033[0m"
	if colors.Apply(input) != expect {
		t.Fail()
	}
}

func TestColorApplyWithBlue(t *testing.T) {
	input := "<blue>hello</blue>"
	expect := "\033[34mhello\033[0m"
	if colors.Apply(input) != expect {
		t.Fail()
	}
}

func TestColorApplyWithUnderline(t *testing.T) {
	input := "<u>hello</u>"
	expect := "\033[4mhello\033[0m"
	if colors.Apply(input) != expect {
		t.Fail()
	}
}

func TestColorApplyWithLight(t *testing.T) {
	input := "<light>hello</light>"
	expect := "\033[1mhello\033[0m"
	if colors.Apply(input) != expect {
		t.Fail()
	}
}

func TestColorApplyWithColorAndUnderline(t *testing.T) {
	input := "<red><u>hello</u></red>"
	expect := "\033[31m\033[4mhello\033[0m\033[0m"
	if colors.Apply(input) != expect {
		t.Fail()
	}
}
