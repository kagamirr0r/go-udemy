package main

import "fmt"

//　関数

func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func Noreturn() {
	fmt.Println("No Return")
	return
}

func main() {
	i := Plus(1,2)
	fmt.Println(i)

	i2, i3 := Div(9, 4)
	fmt.Println(i2, i3)

	i4, _ := Div(9, 3) //返り血を省略できる
	fmt.Println(i4)

	i5 := Double(1000)
	fmt.Println(i5)

	Noreturn()

	// 無名関数
	f := func(x,y int ) int {
		return x + y
	}
	i6 := f(1,2)
	fmt.Println(i6)

	i7 := func(x,y int ) int {
		return x + y
	}(1,2) //ここで渡せる！
	fmt.Println(i7)

	f2 := ReturnFunc()
	f2()

	CallFunction(func() {
		fmt.Println("I'm a Fucntion")
	})


	//クロージャー
	f3 := Later()
	fmt.Println(f3("HELLO"))
	fmt.Println(f3("My"))
	fmt.Println(f3("name"))
	fmt.Println(f3("is"))
	fmt.Println(f3("Go lang"))

	ints := integers()
	fmt.Println(ints()) // 1
	fmt.Println(ints()) // 2
	fmt.Println(ints()) // 3
	fmt.Println(ints()) // 4

	otherints := integers()
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
}

// 関数を返す関数
func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a Function")
	}
}


//関数を引数にとる関数
func CallFunction(f func()) {
	f()
}

// クロージャー
func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

// ジェネレーター
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

