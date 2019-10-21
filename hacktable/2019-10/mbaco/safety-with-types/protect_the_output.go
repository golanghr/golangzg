package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Email    string
	Password string
}

func main() {
	fmt.Println("==== NOT PROTECTED ====")

	u := User{
		Email:    "name.surname@somewhere.com",
		Password: "pass123",
	}

	log.Printf("%+v\n", u)

	data, err := json.Marshal(u)
	if err != nil {
		log.Fatalf("JSON encode failed: %v", err)
	}

	fmt.Printf("%s\n", data)

	fmt.Println("=====================================")

	fmt.Println("==== PROTECTED ====")

	su := SafeUser{
		Email:    "name.surname@somewhere.com",
		Password: "pass123",
	}

	log.Printf("%+v\n", su)

	data, err = json.Marshal(su)
	if err != nil {
		log.Fatalf("JSON encode failed: %v", err)
	}

	fmt.Printf("%s\n", data)

	fmt.Println("=====================================")
}

type SafeUser struct {
	Email    string
	Password password // `json:"-"`
}

type password string

// String satisfy fmt.Stringer.
func (ss password) String() string {
	return "******"
}

// MarshalText satisfy encoding.TextMarshaler.
func (ss password) MarshalText() ([]byte, error) {
	return []byte(`******`), nil
}

// // MarshalJSON satisfy json.Marshaler.
// func (ss password) MarshalJSON() ([]byte, error) {
// 	return []byte(`"******"`), nil
// }
