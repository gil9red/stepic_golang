package main

import (
	"fmt"
	"time"
)

func main() {
	// time.Now | time.Date | time.Unix | Format
	{
		// func Now() Time
		// возвращает текущую дату и время
		now := time.Now()

		// func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
		// возвращает дату и время в соответствии с заданными параметрами: годом, месяцем, днем, временем и пр.
		currentTime := time.Date(
			2020,     // год
			time.May, // месяц
			15,       // день
			10,       // часы
			13,       // минуты
			12,       // секунды
			45,       // наносекунды
			time.UTC, // временная зона
		)

		// func Unix(sec int64, nsec int64) Time
		// возвращает дату и время в соответствии с заданными параметрами: секундами и наносекундами, прошедшими с начала эпохи Unix — 01.01.1970 г.
		// https://ru.wikipedia.org/wiki/Unix-%D0%B2%D1%80%D0%B5%D0%BC%D1%8F
		unixTime := time.Unix(
			150000, // секунды
			1,      // наносекунды
		)

		fmt.Println(now.Format("02-01-2006 15:04:05"))         // 27-10-2020 10:58:52
		fmt.Println(currentTime.Format("02-01-2006 15:04:05")) // 15-05-2020 10:13:12
		fmt.Println(unixTime.Format("02-01-2006 15:04:05"))    // 02-01-1970 22:40:00

		fmt.Println()
	}

	// time.Parse | time.LoadLocation | time.ParseInLocation | Format
	{
		// func Parse(layout, value string) (Time, error)
		// парсит дату и время в строковом представлении
		firstTime, err := time.Parse("2006/01/02 15-04", "2020/05/15 17-45")
		if err != nil {
			panic(err)
		}

		// LoadLocation находит временную зону в справочнике IANA
		// https://www.iana.org/time-zones
		loc, err := time.LoadLocation("Asia/Yekaterinburg")
		if err != nil {
			panic(err)
		}

		// func ParseInLocation(layout, value string, loc *Location) (Time, error)
		// парсит дату и время в строковом представлении с отдельным указанием временной зоны
		secondTime, err := time.ParseInLocation("Jan 2 06 03:04:05pm", "May 15 20 05:45:00pm", loc)
		if err != nil {
			panic(err)
		}

		fmt.Println(firstTime, "|", secondTime)
		fmt.Println(firstTime.UTC(), "|", secondTime.UTC())
		// 2020-05-15 17:45:00 +0000 UTC | 2020-05-15 17:45:00 +0500 +05
		// 2020-05-15 17:45:00 +0000 UTC | 2020-05-15 12:45:00 +0000 UTC

		const dateLayout = "02-01-2006 15:04:05"
		fmt.Println(
			firstTime.Format(dateLayout),
			"|",
			secondTime.Format(dateLayout),
		)
		fmt.Println(
			firstTime.UTC().Format(dateLayout),
			"|",
			secondTime.UTC().Format(dateLayout),
		)
		// 15-05-2020 17:45:00 | 15-05-2020 17:45:00
		// 15-05-2020 17:45:00 | 15-05-2020 12:45:00
	}
}
