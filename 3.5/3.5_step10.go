package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
)

func createCSVBuffer() *bytes.Buffer {
	// Записывать данные, а в дальнейшем читать их мы будем из буфера,
	// но его можно заменить любым другим объектом, удовлетворяющим
	// интерфейсу io.ReadWriter
	buf := bytes.NewBuffer(nil)

	w := csv.NewWriter(buf)

	for i := 1; i <= 3; i++ {
		// Запись данных может производится поэтапно, например в цикле
		val1 := fmt.Sprintf("row%d_col1", i)
		val2 := fmt.Sprintf("row%d_col2", i)
		val3 := fmt.Sprintf("row%d_col3", i)
		if err := w.Write([]string{val1, val2, val3}); err != nil { // Аргументом Write является срез строк
			panic(err)
		}
	}
	w.Flush() // Этот метод приведет к фактической записи данных из буфера csv.Writer в buf

	// Либо данные можно записать за один раз
	_ = w.WriteAll([][]string{ // Аргументом WriteAll является срез срезов строк
		{"row4_col1", "row4_col2", "row4_col3"},
		{"row5_col1", "row5_col2", "row5_col3"},
	})

	return buf
}

func main() {
	{
		buf := createCSVBuffer()

		data, _ := ioutil.ReadAll(buf)
		fmt.Println(strconv.Quote(string(data)))
		// "row1_col1,row1_col2,row1_col3\nrow2_col1,row2_col2,row2_col3\nrow3_col1,row3_col2,row3_col3\nrow4_col1,row4_col2,row4_col3\nrow5_col1,row5_col2,row5_col3\n"

		fmt.Println()
	}

	{
		buf := createCSVBuffer()
		r := csv.NewReader(buf)

		for i := 0; i < 2; i++ {
			// Читать данные мы тоже можем построчно, получая срез строк за каждую итерацию
			row, err := r.Read()
			if err != nil && err != io.EOF { // Здесь тоже нужно учитывать конец файла
				panic(err)
			}
			fmt.Println(row)
		}
		// [row1_col1 row1_col2 row1_col3]
		// [row2_col1 row2_col2 row2_col3]

		// Либо прочитать данные за один раз
		data, err := r.ReadAll()
		if err != nil {
			// Когда мы читаем данные до конца файла io.EOF не возвращается, а служит сигналом к завершению чтения
			panic(err)
		}

		for _, row := range data {
			fmt.Println(row)
		}
		// [row3_col1 row3_col2 row3_col3]
		// [row4_col1 row4_col2 row4_col3]
		// [row5_col1 row5_col2 row5_col3]
	}
}
