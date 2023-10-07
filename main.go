package main

import (
	"fmt"

	model "github.com/Phadoo131/hahalanguage/model"
)

func main(){
	var haha *model.Hahalanguage
	
	testHa := haha.HahaThis(12345)
	testHa2 := haha.HahaThis("Let's have some fun, shall we?")
	testHa3 := haha.HahaThis("h")
	testHa4 := haha.HahaThis([]string{"Convert", "this", "si,", "kri", "kri", "!"})
	testHa5 := haha.FunOrFalse(12345)
	testHa6 := haha.FunOrFalse(12346)
	testHa7 := haha.WhereIsFun("Where is the fun na?")
	testHa8 := haha.WhereIsFun("I don't believe it's a fun in this")
	testHa9 := haha.FunOrFalse('a')
	testHa11 := haha.FunOrFalse("5")
	testHa10 := haha.FunOrFalse([]string{"Hello a", "5", "yes"})


	fmt.Println(testHa)
	fmt.Println(testHa2)
	fmt.Println(testHa3)
	fmt.Println(testHa4)
	fmt.Println(testHa5)
	fmt.Println(testHa6)
	fmt.Println(testHa7)
	fmt.Println(testHa8)
	fmt.Println(testHa9)
	fmt.Println(testHa10)
	fmt.Println(testHa11)
}