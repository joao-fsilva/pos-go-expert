package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fastjson"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var p fastjson.Parser
	jsonData := `{"user":{"name":"Tom","age":25}}`

	v, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	user := v.GetObject("user")

	fmt.Printf("User: %s\n", user)
	fmt.Printf("Name: %s\n", user.Get("name"))
	fmt.Printf("Age: %s\n", user.Get("age"))

	userJSON := v.Get("user").String()
	fmt.Printf("User JSON: %s\n", userJSON)

	var user2 User

	if err := json.Unmarshal([]byte(userJSON), &user2); err != nil {
		panic(err)
	}

	fmt.Println("User Name:", user2.Name)
	fmt.Println("User Age", user2.Age)
}
