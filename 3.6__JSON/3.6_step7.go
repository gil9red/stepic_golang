package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Аннотация выглядит так: `json:"используемое_имя,опция,опция"`

	type myStruct struct {
		// При кодировании / декодировании будет использовано имя name, а не Name
		Name string `json:"name"`

		Comment string

		// При кодировании / декодировании будет использовано то же имя (Age),
		// но если значение поля равно 0 (пустое значение: false, nil, пустой слайс и пр.),
		// то при кодировании оно будет опущено
		Age int `json:",omitempty"`

		// При кодировании / декодировании поле всегда игнорируется
		Status bool `json:"-"`

		// Поля в нижнем регистре не (де)сериализуются
		notExported string
	}

	m := myStruct{Name: "John Connor", Age: 0, Status: true}
	data, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)
	// {"name":"John Connor","Comment":""}

	m = myStruct{Name: "John Connor", Age: 99, Status: false, Comment: "Not alive"}
	data, err = json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)
	// {"name":"John Connor","Comment":"Not alive","Age":99}
}
