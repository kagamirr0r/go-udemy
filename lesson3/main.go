package main

import "fmt"

//変数

// i5 := 500 関数外で定義できない
var i5 = 500 //関数外で定義できる

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	//明示的な定義
	// var 変数名 型 = 値

	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int = 200
		s2 string = "Go"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string

	fmt.Println(i3) // 0
	fmt.Println(s3) // ""

	i3 = 300
	s3 = "Go"

	fmt.Println(i3) //変更されている
	fmt.Println(s3) //変更されている

	i = 150
	fmt.Println(i) //変更されている


	//暗黙的な定義（型推論）
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// i4 := 400 再定義はできない
	// i4 = "Go lang" 型エラーになる

	fmt.Println(i5)
	outer()

	// fmt.Println(s4) //他の関数の変数は使えない
}