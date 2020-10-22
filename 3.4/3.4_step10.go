package main

import (
	"errors"
	"fmt"
)

/*
Пришло время для задач, где вы сможете применить полученные знания на практике.

Обязательные условия выполнения: данные со стандартного ввода читаются функцией readTask(),
которая возвращает 3 значения типа пустой интерфейс. Эта функция использует пакеты
encoding/json, fmt, и os - не удаляйте их из импорта. Скорее всего, вам понадобится
пакет "fmt", но вы можете использовать любой другой пакет для записи в стандартный вывод
не удаляя fmt.

И так, вы получаете 3 значения типа пустой интерфейс: если все удачно, то первые 2 значения
будут float64. Третье значение в идеальном случае будет строкой, которая может иметь
значения: "+", "-", "*", "/" (определенная математическая операция). Но такие идеальные
случаи будут не всегда, вы можете получить значения других типов, а также строка в третьем
значении может не относится к одной из указанных математических операций.

Результат выполнения программы должен быть таков:
    1. в штатной ситуации вы должны напечатать в стандартный вывод результат выполнения
       математической операции с точностью до 4 цифры после запятой (fmt.Printf(%.4f));
    2. если первое или второе значение не является типом float64, вы должны напечатать
       сообщение об ошибке вида: value=полученное_значение: тип_значения
       (например: value=true: bool)
    3. если третье значение имеет неверный тип или передан знак, не относящийся к указанным
       выше математическим операциям, сообщение об ошибке должно иметь вид: неизвестная операция

Гарантируется, что ошибка в аргументах может быть только одна, поэтому если вы при проверке
первого значения увидели, что оно содержит ошибку - печатайте сообщение об ошибке и
завершайте работу программы, проверка остальных аргументов значения уже не имеет, а
проверяющая система воспримет 2 сообщения об ошибке как нарушение условия выполнения задания.

Удачи!
*/

func getValue(value interface{}) (float64, error) {
	if v, ok := value.(float64); ok {
		return v, nil
	}
	return 0, errors.New(fmt.Sprintf("value=%v: %T", value, value))
}

func getOperation(operation interface{}) (fn func(a float64, b float64) float64, err error) {
	v, ok := operation.(string)
	if !ok {
		return nil, errors.New(fmt.Sprintf("invalid type '%T', must be string", operation))
	}

	switch v {
	case "+":
		fn = func(a float64, b float64) float64 {
			return a + b
		}
	case "-":
		fn = func(a float64, b float64) float64 {
			return a - b
		}
	case "*":
		fn = func(a float64, b float64) float64 {
			return a * b
		}
	case "/":
		fn = func(a float64, b float64) float64 {
			return a / b
		}
	default:
		err = errors.New("неизвестная операция")
	}

	return
}

func readTask() (interface{}, interface{}, interface{}) {
	return 5.4, 2.0, "+"
}

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	value1, value2, operation := readTask()

	var num1, num2 float64
	var err error
	var fnOper func(a float64, b float64) float64

	num1, err = getValue(value1)
	if err != nil {
		panic(err)
	}

	num2, err = getValue(value2)
	if err != nil {
		panic(err)
	}

	fnOper, err = getOperation(operation)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%.4f\n", fnOper(num1, num2))
}
