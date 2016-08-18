package search

import (
	"fmt"
	"net/http"
	"time"
)

func Query(url string) *http.Response {
	ch := make(chan *http.Response)
	go func() {
		resp, _ := http.Get(url)
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
