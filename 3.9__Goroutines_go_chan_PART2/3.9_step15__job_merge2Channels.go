package main

import (
	"fmt"
	"time"
)

/*
Необходимо написать функцию
    func merge2Channels(fn func(int) int, in1 <-chan int, in2 <- chan int, out chan<- int, n int).

Описание ее работы:
n раз сделать следующее
 * прочитать по одному числу из каждого из двух каналов in1 и in2, назовем их x1 и x2.
 * вычислить f(x1) + f(x2)
 * записать полученное значение в out

Функция merge2Channels должна быть неблокирующей, сразу возвращая управление.

Функция fn может работать долгое время, ожидая чего-либо или производя вычисления.

Формат ввода:
 * количество итераций передается через аргумент n.
 * целые числа подаются через аргументы-каналы in1 и in2.
 * функция для обработки чисел перед сложением передается через аргумент fn.

Формат вывода:
 * канал для вывода результатов передается через аргумент out.
*/

// SOURCE: https://habr.com/ru/company/ozontech/blog/507932/
func merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	idx := make(chan int) // Для передачи индекса
	val := make(chan int) // Для передачи значения

	go func() {
		for i := 0; i < n; i++ {
			x1 := <-in1
			x2 := <-in2

			done := make(chan int)
			go func() {
				done <- f(x1)
			}()
			go func() {
				done <- f(x2)
			}()

			go func(i int) {
				idx <- i
				val <- (<-done) + (<-done)
			}(i)
		}
	}()

	go func() {
		m := make(map[int]int)
		for i := 0; i < n; i++ {
			m[<-idx] = <-val
		}

		// Для возврата значений по порядку
		for i := 0; i < n; i++ {
			out <- m[i]
		}
	}()
}

func main() {
	fmt.Println("Started!")

	start := time.Now()

	fn := func(x int) int {
		time.Sleep(time.Second * 2)
		return x
	}

	in1 := make(chan int)
	in2 := make(chan int)
	out := make(chan int)
	n := 5

	go func() {
		for i := 0; i < n; i++ {
			fmt.Println("Send", i)
			go func() {
				in1 <- 1
				in2 <- 2
			}()
		}
	}()

	merge2Channels(fn, in1, in2, out, n)

	for i := 0; i < n; i++ {
		value := <-out
		fmt.Println("Result:", value)
	}

	fmt.Println("Elapsed:", time.Since(start))
	fmt.Println("Finish!")
}
