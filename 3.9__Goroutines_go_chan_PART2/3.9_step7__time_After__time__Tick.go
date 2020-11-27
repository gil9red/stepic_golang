package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Started!")

	// func Sleep(d Duration)
	// программа засыпает на заданное время
	time.Sleep(time.Second * 2) // спим ровно 2 секунды
	fmt.Println("Прошло 2 секунды")

	// func After(d Duration) <-chan Time
	// создает канал, который через заданное время вернет значение
	timer := time.After(time.Second)
	<-timer // значение будет получено из канала ровно через 1 секунду
	fmt.Println("Прошла 1 секунда")

	// func Tick(d Duration) <-chan Time
	// создает канал, который будет посылать сигналы постоянно через заданный промежуток времени
	ticker := time.Tick(time.Second)
	count := 0

	for {
		<-ticker
		fmt.Println("Очередной тик")
		count++
		if count == 3 {
			break
		}
	}
	// Очередной тик
	// Очередной тик
	// Очередной тик

	fmt.Println("Finish!")
}
