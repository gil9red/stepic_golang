package main

import (
	"fmt"
	"time"
)

/*
На стандартный ввод подается строковое представление даты и времени в следующем формате:
1986-04-16T05:20:00+06:00

Ваша задача конвертировать эту строку в Time, а затем вывести в формате UnixDate:
Wed Apr 16 05:20:00 +0600 1986

Для более эффективной работы рекомендуется ознакомиться с документацией о содержащихся в
модуле time константах.

Sample Input:
1986-04-16T05:20:00+06:00

Sample Output:
Wed Apr 16 05:20:00 +0600 1986
*/

func main() {
	var dateStr string
	_, err := fmt.Scan(&dateStr)
	if err != nil {
		panic(err)
	}
	//dateStr = "1986-04-16T05:20:00+06:00"

	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		panic(err)
	}

	fmt.Println(date.Format(time.UnixDate))
}
