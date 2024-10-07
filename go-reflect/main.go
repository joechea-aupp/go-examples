package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type City struct {
	Name       string  `json:"name"`
	Population int64   `json:"population"`
	GDP        float64 `json:"gdp"`
	Mayjor     string  `json:"mayjor"`
}

// law of reflection
// 1. you can go from interface value to reflection obj by wrapping the actual value in a reflection with ValueOf and TypeOf
// 1.1 ValueOf is use to wrap value from a variable
// 1.2 TypeOf is use to wrap type from a variable
// the benefit of having it wraped is, we can perform action on top of whatever go provided at runtime.

// 2. we can go from reflection object to interface value, easily by just calling a method

func JSONEncode(v interface{}) ([]byte, error) {
	// wrape value with reflection
	refObjVal := reflect.ValueOf(v)
	// wrape type with reflection
	refObjType := reflect.TypeOf(v)

	buf := bytes.Buffer{}
	pairs := []string{}

	// NumField return total field that struct has
	for i := 0; i < refObjVal.NumField(); i++ {
		// Field provide the ability to reflect on the field level.
		// Field value will get the value action for reflection
		structFieldRefObj := refObjVal.Field(i)
		// Field type will get the type action for reflection
		structFieldRefObjTyp := refObjType.Field(i)

		// struct type is reside within the type refection
		tag := structFieldRefObjTyp.Tag.Get("json")

		// Kind() similar to type, the difference is Kind() only show the basic builtin type from the system (go)
		switch structFieldRefObj.Kind() {
		case reflect.String:
			// We can use .String, or .Int to get the value to the interface however, by using Interface() it will retrieve the original value before being wrape with reflection
			strVal := structFieldRefObj.Interface().(string)
			jsonStringFormat := fmt.Sprintf(`"%s":"%s"`, tag, strVal)
			pairs = append(pairs, jsonStringFormat)

		case reflect.Int64:
			intVal := structFieldRefObj.Interface().(int64)
			jsonIntFormat := fmt.Sprintf(`"%s":%d`, tag, intVal)
			pairs = append(pairs, jsonIntFormat)

		default:
			// .Name when use with type, it will return the field name.
			return buf.Bytes(), fmt.Errorf("struct field with name %s and kind %s is not support",
				structFieldRefObjTyp.Name, structFieldRefObj.Kind())
		}

	}
	buf.WriteString("{")
	buf.WriteString(strings.Join(pairs, ","))
	buf.WriteString("}")

	return buf.Bytes(), nil
}

func main() {
	user := User{
		"Joe", 18,
	}

	data, err := JSONEncode(user)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
