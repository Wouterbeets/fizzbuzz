package http

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestFizzHandler(t *testing.T) {
	client := &http.Client{}
	tests := []struct {
		body string
	}{
		{
			body: "test",
		},
	}
	go NewServer().ListenAndServe()
	for _, t := range tests {
		r := strings.NewReader(t.body)
		resp, err := client.Post("http://localhost:8000", "application/json", r)
		fmt.Printf("%+v, %+v\n", resp, err)
	}
}
