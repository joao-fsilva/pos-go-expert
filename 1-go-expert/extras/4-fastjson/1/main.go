package main

import "github.com/valyala/fastjson"

func main() {
	var p fastjson.Parser

	jsonData := `{"foo": "bar", "arr": [1, 2, 3]}`

	v, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	foo := v.GetStringBytes("foo")
	println(string(foo))

	arr := v.GetArray("arr")
	for i, v := range arr {
		println(arr[i].GetInt())
		println(v.GetInt())
	}
}
