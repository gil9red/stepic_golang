package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	type myStruct struct {
		Name   string
		Age    int
		Status bool
		Values []int
	}
	s := myStruct{
		Name:   "John Connor",
		Age:    35,
		Status: true,
		Values: []int{15, 11, 37},
	}

	// json.Marshal
	{
		// Функция Marshal принимает аргумент типа interface{} (в нашем случае это структура)
		// и возвращает байтовый срез с данными, кодированными в формат JSON.
		data, err := json.Marshal(s)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n", data)
		// {"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}

		fmt.Println()
	}

	// json.MarshalIndent
	{
		// MarshalIndent похож на Marshal, но применяет отступ (indent) для форматирования вывода.
		// Каждый элемент JSON в выходных данных начинается с новой строки, начинающейся с
		// префикса (prefix), за которым следует один или несколько отступов в соответствии с вложенностью:
		data, err := json.MarshalIndent(s, "", "    ")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s\n", data)
		// {
		//     "Name": "John Connor",
		//     "Age": 35,
		//     "Status": true,
		//     "Values": [
		//	     15,
		//	     11,
		//	     37
		//     ]
		// }

		fmt.Println()
	}

	// json.Unmarshal
	{
		var s myStruct

		data := []byte(`{"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}`)
		if err := json.Unmarshal(data, &s); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%v\n", s)
		// {John Connor 35 true [15 11 37]}

		data = []byte(`
			{
				"Name": "John Connor",
				"Age": 35,
				"Status": true,
				"Values": [
					15,
					11,
					37
				]
			}
		`)
		if err := json.Unmarshal(data, &s); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%v\n", s)
		// {John Connor 35 true [15 11 37]}

		fmt.Println()
	}

	// json.Valid
	{
		data, _ := json.Marshal(s)
		data = bytes.Trim(data, "{") // испортим json удалив '{'

		// функция json.Valid возвращает bool, true - если json правильный
		if !json.Valid(data) {
			fmt.Printf("invalid json: %s\n", data)
		}
		// invalid json: "Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}
	}
}
