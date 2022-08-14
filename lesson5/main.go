package main

import "fmt"

//定数

const Pi = 3.14

const (
	URL = "http://xxx.co.jp"
	SiteName = "test"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

// var Big int = 9223372036854775807 + 1 //as int value in variable declaration (overflows)
const Big = 9223372036854775807 + 1 //as int value in variable declaration (overflows)

const (
	c0 = iota
	c1
	c2
)

func main() {
		fmt.Println(Pi)

		// Pi = 3 //cannot assign to Pi

		fmt.Println(URL, SiteName)

		fmt.Println(A,B,C,D,E,F) //1 1 1 A A A
		fmt.Println(Big -1)

		fmt.Println(c0, c1,c2) //0 1 2
}