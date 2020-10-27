package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
На стандартный ввод подаются данные о продолжительности периода (формат приведен в примере). Кроме того,
вам дана дата в формате Unix-Time: 1589570165 в виде константы типа int64 (наносекунды для целей
преобразования равны 0).

Требуется считать данные о продолжительности периода, преобразовать их в тип Duration, а затем
вывести (в формате UnixDate) дату и время, получившиеся при добавлении периода к стандартной дате.

Небольшая подсказка: базовую дату необходимо явно перенести в зону UTC с помощью одноименного метода.

Sample Input:
12 мин. 13 сек.

Sample Output:
Fri May 15 19:28:18 UTC 2020
*/

func main() {
	startDate := time.Unix(1589570165, 0)

	var text string
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	text = strings.TrimSpace(text)

	//text = "12 мин. 13 сек."
	//text = "71 мин. 8 сек."

	re := regexp.MustCompile(`(\d+) мин\. (\d+) сек\.`)
	items := re.FindStringSubmatch(text)
	if len(items) == 0 {
		panic("Invalid text format!")
	}

	minutes, _ := strconv.Atoi(items[1])
	seconds, _ := strconv.Atoi(items[2])
	d := time.Minute*time.Duration(minutes) + time.Second*time.Duration(seconds)
	fmt.Println(startDate.Add(d).UTC().Format(time.UnixDate))
}
