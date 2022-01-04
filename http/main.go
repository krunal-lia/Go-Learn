package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://qtalk.in/getMetaFromUrl/v2/https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3D2UVbOjqrWf0m")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}
	io.Copy(lw, resp.Body)
	// body, error := io.ReadAll(resp.Body)

	// if error != nil {
	// 	fmt.Println("Error", error)
	// }
	// fmt.Println(string(body))
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
