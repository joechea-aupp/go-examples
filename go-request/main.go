package main

import (
	"fmt"
	"net/url"
	"reflect"
)

import (
	"net/http"
)

type User struct {
	Name      string `form:"[user]name"`
	ShortName string `form:"[user]short_name"`
}

type Pseudonym struct {
	UniqueID                 string `form:"[pseudonym]unique_id"`
	Password                 string `form:"[pseudonym]password"`
	SISUserID                string `form:"[pseudonym]sis_user_id"`
	AuthenticationProviderID string `form:"[pseudonym]authentication_provider_id"`
}

type UserRegistration struct {
	User      User      `form:"[user]"`
	Pseudonym Pseudonym `form:"[pseudonym]"`
}

func encodeStructToNestedForm(data interface{}) url.Values {
	values := url.Values{}
	v := reflect.ValueOf(data)

	fmt.Println("value for v:", v.Field(0))
	fmt.Println("type for v:", v.Type().Field(0).Name)

	var encodeField func(reflect.Value, string)
	encodeField = func(v reflect.Value, prefix string) {
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			value := v.Field(i)

			if value.Kind() == reflect.Struct {
				encodeField(value, prefix+field.Name+".")
			} else {
				tag := field.Tag.Get("form")
				if tag != "" {
					values.Set(tag, fmt.Sprintf("%v", value.Interface()))
				}
			}
		}
	}

	encodeField(v, "")
	return values
}

func main() {
	data := UserRegistration{
		User: User{
			Name:      "John Doe",
			ShortName: "JD",
		},
		Pseudonym: Pseudonym{
			UniqueID:                 "john.doe",
			Password:                 "password",
			SISUserID:                "123456",
			AuthenticationProviderID: "saml",
		},
	}
	encodeData := encodeStructToNestedForm(data)
	fmt.Println(encodeData.Encode())

	// client := &http.Client{}

	http.PostForm("http://example.com", encodeData)
}
