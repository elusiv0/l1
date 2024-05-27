package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "abAsBs"
	str2 := "aBnBoOiii"
	str3 := "ABCVDoiuy"
	str4 := "lkjhgfd"
	fmt.Printf("%v %v %v %v", checkUniq(str1), checkUniq(str2), checkUniq(str3), checkUniq(str4))
}

func checkUniq(str string) bool {
	str = strings.ToLower(str)
	set := make(map[rune]bool) //заводим set встреченных эл-тв

	for _, elem := range str {
		if set[elem] { //если был - возвращаем ложно
			return false
		}
		set[elem] = true //добавляем информацию о встреченной переменной
	}

	return true //возврашаем истинно, если все эл-ты уникальны
}
