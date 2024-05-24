package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Second}
	//	resp, err := c.Get("http://www.google.com")
	jsonVar := bytes.NewBuffer([]byte(`{"name":"John"}`))
	resp, err := c.Post("http://www.google.com", "application/json", jsonVar)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
