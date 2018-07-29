package parser_test

import (
	"testing"

	"github.com/huytd/stackoverlazy/parser"
	"net/http"
	"net/http/httptest"
	"io"
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

func TestParseURLParseHttpsUrl(t *testing.T) {
	expect := parser.ParseAnswer(fakeHtml())
	if expect == "" {
		t.Fail()
	}
}

func fakeHtml() *http.Response {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body><a href=\"https://stackoverflow.com\"></a><a href=\"http://stackoverflow.com\"></a></body></html>")
	}

	req := httptest.NewRequest("GET", "http://google.com/search?q=stackoverflow+sample", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	return w.Result()
}
