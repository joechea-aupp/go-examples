package main

import (
	"bytes"
	"encoding/json"
	"os"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	f, err := os.OpenFile("file.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// instanciate new buffer point to buf variable
	buf := new(bytes.Buffer)

	// create json encoder that write to the newly create buffer instance
	// the reason why buf can be use with NewEncoder because bytes.Buffer implement io.Writer interface (Writer)
	enc := json.NewEncoder(buf)
	u := user{"Jack", 18}
	// encode data to buffer
	if err := enc.Encode(u); err != nil {
		panic(err)
	}
	// write buffer data to a file
	f.Write(buf.Bytes())
}
