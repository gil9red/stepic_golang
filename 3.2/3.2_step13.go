package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
Представьте что вы работаете в большой компании где используется
модульная архитектура. Ваш коллега написал модуль с какой-то логикой
(вы не знаете) и передает в вашу программу какие-то данные. Вы же пишите
функцию которая считывает две переменных типа string, а возвращает число
типа int64 которое нужно получить сложением этим строк.

Но было бы так все просто, ведь ваш коллега не пишет на Go, и он зол из-за
того что гоферам платят больше. Поэтому он решил подшутить над вами и подсунул
вам подвох. Он придумал вставлять мусор в строки перед тем как вызывать вашу функцию.

Поэтому предварительно вам нужно убрать из них мусор и конвертировать в числа.
Под мусором имеются ввиду лишние символы и спец знаки. Разрешается использовать
только определенные пакеты: fmt, strconv, unicode, strings, bytes.

Sample Input:
%^80 hhhhh20&&&&nd

Sample Output:
100
*/

func ParseNumber(numStr string) int64 {
	builder := strings.Builder{}
	for _, r := range numStr {
		if unicode.Is(unicode.Digit, r) {
			builder.WriteRune(r)
		}
	}

	number, err := strconv.ParseInt(builder.String(), 10, 64)
	if err != nil {
		panic(err)
	}

	return number
}

func adding(a, b string) int64 {
	return ParseNumber(a) + ParseNumber(b)
}

func main() {
	fmt.Println(adding("%^80", "hhhhh20&&&&nd"))
}
