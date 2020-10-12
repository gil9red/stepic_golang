package main

import "fmt"

/*
Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power,
с типами bool, int, int соответственно. У этой структуры должны быть методы:
Shoot и RideBike, которые не принимают аргументов, но возвращают значение bool.
Если значение On == false, то оба метода вернут false.

Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу,
а метод возвращает true), если его нет, то метод вернет false.
Метод RideBike работает также, но только зависит от свойства Power.

Чтобы проверить, что вы все сделали правильно, вы должны создать указатель на
экземпляр этой структуры с именем testStruct в функции main, в дальнейшем
программа проверит результат.
*/

type Terminator struct {
	On          bool
	Ammo, Power uint
}

func (obj *Terminator) Shoot() bool {
	if obj.On == false || obj.Ammo == 0 {
		return false
	}
	obj.Ammo--
	return true
}

func (obj *Terminator) RideBike() bool {
	if obj.On == false || obj.Power == 0 {
		return false
	}
	obj.Power--
	return true
}

func main() {
	testStruct := new(Terminator)
	fmt.Println(testStruct)
}
