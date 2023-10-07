package model

import (
	"strconv"
	"strings"
	"unicode"

)

func (haha *Hahalanguage) HahahaInTH() int{
	return 555
}

func (haha *Hahalanguage) HahaThis(T any) string{
	stringHa := ""

	if T == nil{
		return "Hahaha it's null na, lol"
	}

	if val, ok := T.(int); ok{
		stringJa := strconv.Itoa(val)
		n := len(stringJa)
		
		for i := 0; i < n; i++{
			stringHa += "5"
		}
	}

	if _, ok := T.(byte); ok{
		return "5"
	}

	if word, ok := T.(string); ok{
		splitThis := strings.Split(word, " ")
		n := len(splitThis)
		for i := 0; i < n; i++{
			nn := len(splitThis[i])
			for j := 0; j < nn; j++{
				char := splitThis[i][j]
				charMap := make(map[byte]int)

				if unicode.IsLetter(rune(char)){
					charMap[char - 'a']++
				}
	
				if _, ok := charMap[char - 'a']; !ok{
					stringHa += " lol "
				} else{
					stringHa += "ha"
				}
			}
		}
	}
	if _, ok := T.(rune); ok{
		return "5"
	}

	if words, ok := T.([]string); ok{
		n := len(words)
		for i := 0; i < n; i++{
			nn := len(words[i])
			for j := 0; j < nn; j++{
				char := words[i][j]
				charMap := make(map[byte]int)

				if unicode.IsLetter(rune(char)){
					charMap[char - 'a']++
				}
	
				if _, ok := charMap[char - 'a']; !ok{
					stringHa += " lol "
				} else{
					stringHa += "ha"
				}
			}
		}
	}
	if words, ok := T.([]int); ok{
		n := len(words)
		for i := 0; i < n; i++{
			stringHa += "5"
		}
	}
	if words, ok := T.([]rune); ok{
		n := len(words)
		for i := 0; i < n; i++{
			stringHa += "5"
		}
	}
	if words, ok := T.([]byte); ok{
		n := len(words)
		for i := 0; i < n; i++{
			stringHa += "ha"
		}
	}


	return stringHa
}