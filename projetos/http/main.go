package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://www.google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// 1 - First approach to print page body

	// fmt.Println("1 - First approach")
	// fmt.Print("\n\n")

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)

	// fmt.Println(string(bs))

	fmt.Print("\n\n")
	fmt.Println("2 - Second approach")
	fmt.Print("\n\n")

	// 2 - Second approach to print page body

	// io.Copy(os.Stdout, resp.Body)

	fmt.Print("\n\n\n\n")

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Wrote bytes:", len(bs))
	return len(bs), nil
}
