package model

import (
	"strconv"
	"strings"
)

func (haha *Hahalanguage) FunOrFalse(T any) bool {
    switch val := T.(type) {
    case int:
        stringJa := strconv.Itoa(val)
        n := len(stringJa)

        for i := 0; i < n; i++ {
            if string(stringJa[i]) == "5" {
                return true
            }
        }
    case rune:
        if string(val) == "5" {
            return true
        }
    case byte:
        if string(val) == "5" {
            return true
        }
    case string:
        n := len(val)
        val = strings.ToLower(val)

        for i := 0; i < n-1; i++ {
            char := string(val[i])
            if char == "5" {
                return true
            } else if char == "h" {
                if string(val[i+1]) == "a" {
                    return true
                }
            } else if char == "l" {
                if string(val[i+1]) == "o" && string(val[i+2]) == "l" {
                    return true
                }
            }
        }
    case []string:
        if len(val) == 0 {
            return false
        }

        for _, s := range val {
            n := len(s)
            for i := 0; i < n; i++ {
                char := string(s[i])
                if char == "5" {
                    return true
                } else if char == "h" {
                    if i < n-1 && string(s[i+1]) == "a" {
                        return true
                    }
                } else if char == "l" {
                    if i < n-2 && string(s[i+1]) == "o" && string(s[i+2]) == "l" {
                        return true
                    }
                }
            }
        }
    case []int:
        if len(val) == 0 {
            return false
        }

        for _, num := range val {
            stringInt := strconv.Itoa(num)
            n := len(stringInt)
            for i := 0; i < n; i++ {
                char := string(stringInt[i])
                if char == "5" {
                    return true
                }
            }
        }
    case []rune:
        return false
    case []byte:
        return false
    }
    return false
}


func (haha *Hahalanguage) WhereIsFun(text string) string {
    preResult := "The fun is at the index: "
    result := ""

    n := len(text)
    lowerText := strings.ToLower(text)
    storeH := -1 
    storeA := -1 

    for i := 0; i < n; i++ {
        if string(lowerText[i]) == "h" && storeH == -1 {
            storeH = i
        }
        if string(lowerText[i]) == "a" && storeA < storeH {
            storeA = i
        }
    }

    if storeH != -1 && storeA != -1 && storeH < storeA {
        result += strconv.Itoa(storeH) + ", " + strconv.Itoa(storeA)
    } else {
        return preResult + "No, there's no fun in this"
    }

    return preResult + result
}