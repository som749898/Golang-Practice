package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.google.com/")

	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))

}
