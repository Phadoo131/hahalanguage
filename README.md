# Hahalanguage

Hahalanguage is a Go programming language package that provides fun and "haha" related functionality. It includes functions to check for the presence of "5" or "ha" in various data types and a function to locate the positions of "h" and "a" in a given string.

Let's have some fun!

## Installation

To use this package, you can include it in your Go project using the following import statement:

```go
import "github.com/Phadoo131/hahalanguage/model"
```

Functions:

FunOrFalse(T any) bool
This function checks if the provided input contains the element "5" or "ha" and returns true if found, otherwise false. It supports various data types including integers, runes, bytes, strings, slices of strings, and slices of integers.

WhereIsFun(text string) string
This function searches for the positions of "h" and "a" in a given string. If "h" is found before "a," it returns the positions as a string. Otherwise, it returns a message indicating that there's no fun in the input string.

HahahaInTH() int
A fun function that always returns 555. Because why not?

HahaThis(T any) string
This function transforms the provided input into a string containing "5" or "ha" or "lol" based on its data type. It's all about fun and "haha."

Usage
Here's an example of how to use the Hahalanguage package:

```go
package main

import (
    "fmt"
    "github.com/Phadoo131/hahalanguage/model"
)

func main() {
    haha := &model.Hahalanguage{}

    // Check if "5" or "ha" is present in various data types
    result := haha.FunOrFalse(12345)
    fmt.Println(result) // true

    // Find the positions of "h" and "a" in a string
    text := "Have some fun!"
    position := haha.WhereIsFun(text)
    fmt.Println(position) // "The fun is at the index: 0, 1"

    // Have some fun with HahaThis
    funString := haha.HahaThis("hello world")
    fmt.Println(funString) // "hahaha lol lol lol lol lololololol"

    // Get your daily dose of "5"
    fiveString := haha.HahaThis([]int{1, 2, 3})
    fmt.Println(fiveString) // "555"

    // Enjoy endless "ha"s
    haString := haha.HahaThis([]byte{1, 2, 3})
    fmt.Println(haString) // "hahaha"
}
```
