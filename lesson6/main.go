package main

import "fmt"

func main() {
	fmt.Println(1 + 2)
	fmt.Println("ABC" + "DE")

	fmt.Println(5 - 1)

	fmt.Println(5 * 4)

	fmt.Println(60 / 3)

	fmt.Println(9 % 4)

	n := 1
	n += 4
	fmt.Println(n)

	n++
	fmt.Println(n)

	n--
	fmt.Println(n)

	s := "ABC"
	s += "DEF"
	fmt.Println(s)

	comparison_operator()
	logical_operator()
}

func comparison_operator() {
		fmt.Println(1 == 1) //true
		fmt.Println(1 == 2) //false

		fmt.Println(4 <= 8) //true
		fmt.Println(1 >= 8) //false

		fmt.Println(4 < 8) //true
		fmt.Println(1 > 8) //false


		fmt.Println(true == false) // false
		fmt.Println(true != false) // true
}

func logical_operator() {
		fmt.Println(true && false == true) //false
		fmt.Println(true && true == true) // true
		fmt.Println(true && false == false) //false

		fmt.Println(true || false == true) //true
		fmt.Println(!true) //false
		fmt.Println(!false) //false
}