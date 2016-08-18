package search

import (
	"fmt"
	"net/http"
	"time"
)

func Query(query string) *http.Response {
	ch := make(chan *http.Response)
	go func() {
		fmt.Print("Looking for answer")
		resp, _ := http.Get("https://search.yahoo.com/search?p=" + query)
		resp.Body.Close()
		ch <- resp
	}()
	for {
		select {
		case response := <-ch:
			return response
		case <-time.After(50 * time.Millisecond):
			fmt.Print(".")
		}
	}
	return nil
}
