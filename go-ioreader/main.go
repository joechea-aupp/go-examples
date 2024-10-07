package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	connReader()
}

func connReader() {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Fprint(conn, "GET HTTP 1.0\r\n\r\n")

	readerToStdout(conn, 25)
}

func stringsReader() {
	s := strings.NewReader("Hello world!")
	readerToStdout(s, 3)
}

func fileReader() {
	f, err := os.Open("ioreader")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	readerToStdout(f, 20)
}

func readerToStdout(r io.Reader, bufSize int) {
	buf := make([]byte, bufSize)
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		if n > 0 {
			fmt.Println(string(buf[:n]))
		}
	}
}
