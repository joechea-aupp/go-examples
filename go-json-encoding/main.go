package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type City struct {
	Name string `json:"city_name"`
	GDP  int    `json:"city_gdp"`
	// GDP        int    `json:"-"` to exclude this field from json output
	Population int `json:"city_pop"`
}

type User struct {
	Name      string      `json:"name"`
	Age       int         `json:"age"`
	City      City        `json:"city"`
	CreatedAt customTime  `json:"created_at"`
	DeletedAt *customTime `json:"deleted_at,omitempty"`
}

type customTime struct {
	time.Time
}

const layout = "2006-01-02"

func (c customTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", c.Format(layout))), nil
}

func (c *customTime) UnmarshalJSON(v []byte) error {
	var err error
	c.Time, err = time.Parse(layout, strings.ReplaceAll(string(v), "\"", ""))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// Marshal -> from struct to json (byte slice)
	fmt.Println("encode from struct to json using: marshal")
	u := User{
		Name: "bob",
		Age:  20,
		City: City{
			Name:       "Phnom Penh",
			GDP:        3177552244,
			Population: 177551222,
		},
		CreatedAt: customTime{
			time.Now(),
		},
	}

	out, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	// Unmarshal -> from json to go struct
	fmt.Println("decode from json to struct using: unmarshal")
	jsonBytes, err := os.Open("out.json")
	if err != nil {
		panic(err)
	}

	defer jsonBytes.Close()

	data := &User{}

	encodeData, err := io.ReadAll(jsonBytes)

	if err := json.Unmarshal(encodeData, data); err != nil {
		panic(err)
	}

	fmt.Println(data)

	// Using json decoder, because we dont want to store []byte in a memory
	fmt.Println("decode from json to struct using: decoder")
	jsonBytes, err = os.Open("out.json")
	if err != nil {
		panic(err)
	}

	data = &User{}

	if err := json.NewDecoder(jsonBytes).Decode(data); err != nil {
		panic(err)
	}

	fmt.Println(data)

	// using json encoder
	fmt.Println("encode from struct to json using: encoder")
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(u); err != nil {
		panic(err)
	}

	fmt.Println(string(buf.Bytes()))
}
