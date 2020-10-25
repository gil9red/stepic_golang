package main

import (
	"bufio"
	"os"
	"strconv"
)

/*
Задача состоит в следующем: на стандартный ввод подаются целые числа в диапазоне 0-100, каждое
число подается на стандартный ввод с новой строки (количество чисел не известно).
Требуется прочитать все эти числа и вывести в стандартный вывод их сумму.

Несколько подсказок: для чтения вы можете использовать типы bufio.Reader и bufio.Scanner,
а для записи - bufio.Writer. При чтении помните об ошибке io.EOF. Конвертирование данных
из строки в целое число и обратно осуществляется функциями Atoi и Itoa из пакета strconv
соответственно. Чтение производится из стандартного ввода (os.Stdin), а запись - в стандартный
вывод (os.Stdout).

Все указанные в тексте задачи пакеты (strconv, bufio, os, io), уже импортированы (другие
использовать нельзя), package main объявлен.

Sample Input:
33
47
12
79
15

Sample Output:
186
*/

func main() {
	in := bufio.NewScanner(os.Stdin)

	num := 0
	for in.Scan() {
		n, err := strconv.Atoi(in.Text())
		if err != nil {
			panic(err)
		}

		num += n
	}

	out := bufio.NewWriter(os.Stdout)
	_, _ = out.WriteString(strconv.Itoa(num))
	_ = out.Flush()
}
