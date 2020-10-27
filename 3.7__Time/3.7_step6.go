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

		var d time.Duration

		// func Since(t Time) Duration
		// вычисляет период между текущим моментом и заданным временем в прошлом
		d = time.Since(post)
		fmt.Println(d, d.Seconds(), d.Round(time.Second), d.Round(time.Second).Seconds(), d.Round(time.Second).Hours())
		// 24h0m0.0080068s 86400.0080068 24h0m0s 86400 24

		// func Until(t Time) Duration
		// вычисляет период между текущим моментом и заданным временем в будущем
		d = time.Until(future)
		fmt.Println(d, d.Seconds(), d.Round(time.Second), d.Round(time.Second).Seconds(), d.Round(time.Second).Hours())
		// 23h59m59.9919935s 86399.9919935 24h0m0s 86400 24

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
