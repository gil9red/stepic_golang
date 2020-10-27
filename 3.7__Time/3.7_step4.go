package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

/*
На стандартный ввод подается строковое представление даты и времени определенного события в
следующем формате:
2020-05-15 08:00:00

Если время события до обеда (13-00), то ничего менять не нужно, достаточно вывести дату на
стандартный вывод в том же формате.
Если же событие должно произойти после обеда, необходимо перенести его на то же время на
следующий день, а затем вывести на стандартный вывод в том же формате.

Sample Input:
2020-05-15 08:00:00

Sample Output:
2020-05-15 08:00:00
*/

func main() {
	var dateStr string
	dateStr, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}
	dateStr = strings.TrimSpace(dateStr)

	//dateStr = "2020-05-15 08:00:00"
	//dateStr = "2020-05-15 14:00:00"

	const dateLayout = "2006-01-02 15:04:05"

	date, err := time.Parse(dateLayout, dateStr)
	if err != nil {
		panic(err)
	}

	if date.Hour() >= 13 {
		date = date.Add(time.Hour * 24)
	}

	fmt.Println(date.Format(dateLayout))
}
