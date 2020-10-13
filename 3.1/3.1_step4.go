package main

import (
	"fmt"
)

/*
https://stepik.org/lesson/345543/step/4?unit=329288
*/

func main() {
	var s1 []map[int]int
	fmt.Println(s1) // []
	s1 = append(s1, map[int]int{
		1: 123,
		2: 456,
		3: 789,
	})
	s1 = append(s1, map[int]int{
		1: 1,
		2: 4,
		3: 9,
	})
	fmt.Println(s1)                 // [map[1:123 2:456 3:789] map[1:1 2:4 3:9]]
	fmt.Println(s1[0][1], s1[1][3]) // 123 9

	fmt.Println()

	var s2 []map[int]string
	fmt.Println(s2) // []
	s2 = append(s2, map[int]string{
		1: "123",
		2: "456",
		3: "789",
	})
	s2 = append(s2, map[int]string{
		1: "1",
		2: "4",
		3: "9",
	})
	fmt.Println(s2)                 // [map[1:123 2:456 3:789] map[1:1 2:4 3:9]]
	fmt.Println(s2[0][1], s2[1][3]) // 123 9

	fmt.Println()

	m1 := make(map[float32]int)
	m1[1.23] = 123
	m1[4.56] = 456
	m1[7.89] = 789
	fmt.Println(m1) // map[1.23:123 4.56:456 7.89:789]

	fmt.Println()

	m2 := map[string][]string{
		"result_1": {"1", "abc"},
		"result_2": {"123"},
	}
	m2["result_3"] = []string{"999"}
	fmt.Println(m2) // map[result_1:[1 abc] result_2:[123] result_3:[999]]

	fmt.Println()

	m3 := map[string]map[string]string{
		"result_1": {
			"1":   "123",
			"abc": "999",
		},
		"result_2": {
			"value": "123",
		},
	}
	m3["result_3"] = map[string]string{
		"999": "777",
	}
	m3["result_4"] = nil
	fmt.Println(m3) // map[result_1:map[1:123 abc:999] result_2:map[value:123] result_3:map[999:777] result_4:map[]]
}
