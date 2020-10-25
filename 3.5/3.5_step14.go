package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
)

/*
Данная задача в основном ориентирована на изучение типа bufio.Reader, поскольку этот тип позволяет
считывать данные постепенно.

В тестовом файле, который вы можете скачать из нашего репозитория на github.com, содержится
длинный ряд чисел, разделенных символом ";". Требуется найти, на какой позиции находится
число 0 и указать его в качестве ответа. Требуется вывести именно позицию числа, а не индекс
(то-есть порядковый номер, нумерация с 1).

Например:  12;234;6;0;78 :
Правильный ответ будет порядковый номер числа: 4
*/

func main() {
	const URL = "https://github.com/semyon-dev/stepik-go/raw/master/work_with_files/task_sep_1/task.data"

	rs, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer rs.Body.Close()

	rd := bufio.NewReader(rs.Body)

	i := 1
	for {
		s, err := rd.ReadString(';')
		if err == io.EOF {
			break
		}

		if s == "0;" {
			fmt.Println(i)
			break
		}
		i++
	}
}
