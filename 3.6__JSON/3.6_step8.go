package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	type testStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var (
		src = testStruct{Name: "John Connor", Age: 35} // Структура с данными
		dst = testStruct{}                             // Структура без данных
		buf = new(bytes.Buffer)                        // Буфер для чтения и записи
	)

	enc := json.NewEncoder(buf) // Работа с одним и тем же буфером
	dec := json.NewDecoder(buf) // Работа с одним и тем же буфером

	err := enc.Encode(src) // Записываем структуру src в буфер
	if err != nil {
		panic(err)
	}

	err = dec.Decode(&dst) // Из буфера заполняем структуру dst
	if err != nil {
		panic(err)
	}

	fmt.Println(dst)
	// {John Connor 35}
}
