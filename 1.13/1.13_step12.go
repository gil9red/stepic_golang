package main

import "fmt"

/*
По данному числу n закончите фразу "На лугу пасется..." одним из возможных
продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".

Входные данные
Дано число n (0<n<100).

Выходные данные
Программа должна вывести введенное число n и одно из слов (на латинице):
korov, korova или korovy, например, 1 korova, 2 korovy, 5 korov.
Между числом и словом должен стоять ровно один пробел.

Sample Input:
10

Sample Output:
10 korov
*/

func main() {
	var n int
	fmt.Scan(&n)

	n10, n100 := n%10, n%100

	var postFix string
	if n10 == 1 && n100 != 11 {
		postFix = "a"
	} else if (n10 == 2 && n100 != 12) || (n10 == 3 && n100 != 13) || (n10 == 4 && n100 != 14) {
		postFix = "y"
	}

	fmt.Printf("%d korov%s\n", n, postFix)
}
