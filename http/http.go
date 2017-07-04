// http contains all http related code
package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wouterbeets/fizzbuzz/fizz"
	"net/http"
	"net/url"
	"strconv"
)

// New server creates and returns a http server configured to listen to the provided address
func NewServer(addr string) *http.Server {
	http.HandleFunc("/fizz", FizzHandler)
	return &http.Server{
		Addr: addr,
	}
}

// Fizzhandler handles all fizz requests
func FizzHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		writeError(w, http.StatusBadRequest, errors.New("Only GET is Allowed"))
		return
	}

	required := []string{
		"string1",
		"string2",
		"int1",
		"int2",
		"limit",
	}
	params := req.URL.Query()
	if err := checkParameters(params, required); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	fb, err := paramsToFizzBuzz(params, required)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(fb.Generate())
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
	}
}

func writeError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, err)
}

func checkParameters(params url.Values, required []string) error {
	for _, key := range required {
		if _, ok := params[key]; !ok {
			return errors.New("Required parameter " + key + " not present")
		}
		if len(params[key]) != 1 {
			return errors.New("Too many values for parameter: " + key)
		}
	}
	for key := range params {
		if !contains(key, required) {
			return errors.New("Unrecognised paramameter: " + key)
		}
	}
	return nil
}

func contains(s string, sl []string) bool {
	for _, str := range sl {
		if s == str {
			return true
		}
	}
	return false
}

func paramsToFizzBuzz(params url.Values, required []string) (*fizz.FizzBuzz, error) {
	f := &fizz.FizzBuzz{
		FizzName: params.Get(required[0]),
		BuzzName: params.Get(required[1]),
	}
	fbl := []int{}
	for i := 2; i < len(required); i++ {
		val, err := strconv.Atoi(params.Get(required[i]))
		if err != nil {
			return nil, errors.New("cannot convert " + params.Get(required[i]) + " to integer")
		}
		fbl = append(fbl, val)
	}
	f.Fizz = fbl[0]
	f.Buzz = fbl[1]
	f.Limit = fbl[2]
	return f, nil
}
