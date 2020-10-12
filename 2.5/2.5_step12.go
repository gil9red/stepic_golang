package main

import (
	"fmt"
	"unicode"
)

/*
Ваша задача сделать проверку подходит ли пароль вводимый пользователем под
заданные требования. Длина пароля должна быть не менее 5 символов, он должен
содержать только цифры и буквы латинского алфавита. На вход подается строка-пароль.
Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"

Sample Input:
fdsghdfgjsdDD1

Sample Output:
Ok
*/

func IsValidPassword(password string) bool {
	rs := []rune(password)
	if len(rs) < 5 {
		return false
	}

	for _, r := range rs {
		if !unicode.Is(unicode.Latin, r) && !unicode.Is(unicode.Digit, r) {
			return false
		}
	}

	return true
}

func main() {
	var text string
	_, _ = fmt.Scan(&text)

	if IsValidPassword(text) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
