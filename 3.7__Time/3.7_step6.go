package main

import (
	"fmt"
	"time"
)

func main() {
	// Duration
	{
		now := time.Now()
		post := now.AddDate(0, 0, -1)
		future := now.AddDate(0, 0, 1)

		// func Since(t Time) Duration
		// вычисляет период между текущим моментом и заданным временем в прошлом
		fmt.Println(time.Since(post).Round(time.Second))
		// 24h0m0s

		// func Until(t Time) Duration
		// вычисляет период между текущим моментом и заданным временем в будущем
		fmt.Println(time.Until(future).Round(time.Second))
		// 24h0m0s

		// func ParseDuration(s string) (Duration, error)
		// преобразует строку в Duration с использованием аннотаций:
		// "ns" - наносекунды,
		// "us" - микросекунды,
		// "ms" - миллисекунды,
		// "s" - секунды,
		// "m" - минуты,
		// "h" - часы.
		dur, err := time.ParseDuration("1h12m3s")
		if err != nil {
			panic(err)
		}
		fmt.Println(dur, dur.Round(time.Hour), dur.Round(time.Hour).Hours())
		// 1h12m3s 1h0m0s 1

		fmt.Println()
	}

	{
		fmt.Println(time.Hour)        // 1h0m0s
		fmt.Println(time.Minute)      // 1m0s
		fmt.Println(time.Second)      // 1s
		fmt.Println(time.Millisecond) // 1ms
		fmt.Println(time.Microsecond) // 1µs
		fmt.Println(time.Nanosecond)  // 1ns

		fmt.Println()

		d := time.Hour*3 + time.Minute*4 + time.Second*5
		fmt.Println(d, d.Hours(), d.Minutes(), d.Seconds(), d.Milliseconds())
		// 3h4m5s 3.0680555555555555 184.08333333333334 11045 11045000

		fmt.Println(d, d.Round(time.Hour), d.Round(time.Minute), d.Round(time.Second), d.Round(time.Millisecond))
		// 3h4m5s 3h0m0s 3h4m0s 3h4m5s 3h4m5s
	}
}
