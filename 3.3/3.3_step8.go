package main

import (
	"errors"
	"fmt"
)

func someFuncWithPanic(needPanic bool) (err error) {
	defer func() {
		// Отложенный вызов анонимной функции, проверяющей, что работа функции завершена
		// без ошибок. Если функция recover() возвращает что угодно кроме nil, значит в ходе
		// выполнения функции возникла паника.
		if e := recover(); e != nil {
			// Здесь происходит приведение интерфейса (об этом мы расскажем буквально в
			// следующем уроке. Результат приведения присваивается переменной err типа error
			// которая уже объявлена при самом вызове функции someFuncWithPanic.
			err = e.(error)

			// После этого анонимная функция завершает свою работу, паника обработана,
			// переменная err, в которой содержится информации о возникшей панике,
			// возвращается как результат выполнения функции.
		}
	}()

	if needPanic {
		panic(errors.New("fatal error"))
	}

	return err
}

func someFuncWithPanic2() {
	panic(errors.New("fatal error someFuncWithPanic2"))
}

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Recover from:", e)
		}
	}()

	defer func() { fmt.Println(1) }()
	defer func() { fmt.Println(2) }()
	defer func() { fmt.Println(3) }()

	fmt.Println("Err1:", someFuncWithPanic(false))
	fmt.Println("Err2:", someFuncWithPanic(true))

	if err := someFuncWithPanic(false); err != nil {
		fmt.Println("Err3:", err)
	}

	//panic("error 123")
	someFuncWithPanic2()
}
