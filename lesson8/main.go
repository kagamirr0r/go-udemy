package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a := 0
	if a == 2 {
			fmt.Println("two")
	} else if a ==1 {
			fmt.Println("one")
	} else {
			fmt.Println("I don't Know")
	}

	if b := 100; b == 100 { // 変数を定義した後にif分をかける
		fmt.Println("one hundred")
	}

	x := 0

	if x := 2; x == 2 {
		fmt.Println(x) // 2 xのスコープはif文の中だけ
	}

	fmt.Println(x) // 0　xのスコープはif文の中だけ

	error_hundring()
	for_func()
	switch_func()
	type_assetion()
	nested_for()
	defer_function()
	panipani()
	go_routin()
	go_init()
}

func error_hundring() {
	var s string = "100"

	i, err1 := strconv.Atoi(s)

	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Printf("i = %T\n", i)


	var s2 string = "a"
	i2, err2 := strconv.Atoi(s2)

	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Printf("i2 = %T\n", i2)
}

func for_func() {
	// i := 0
	// for {
	// 	i++
	// 	if i == 3 {
	// 		break
	// 	}
	// 	fmt.Println("Loop")
	// }

	// point := 0
	// for point < 10 {
	// 	fmt.Println(point)
	// 	point++
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 3 {
	// 		continue //　３が飛ばされる
	// 	}

	// 	if i == 6 {
	// 		break //　6で終了
	// 	}
	// 	fmt.Println(i)
	// }

	// arr := [3]int{1,2,3}
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// arr := [3]int{1,2,3}
	// for i, v := range arr {
	// 	fmt.Println(i, v)
	// }

	// for _, v := range arr {
	// 	fmt.Println(v)
	// }

	// sl := []string{"Python", "PHP", "Go"}
	// for i, v := range sl {
	// 	fmt.Println(i, v)
	// }


	// fruits := []string{"apple", "banana", "peach"}
	// for i, v := range fruits {
	// 	fmt.Println(i, v)
	// }

	// m := map[string]int{"apple": 100, "banana": 150}
	// for k,v := range m {
	// 	fmt.Println(k, v)
	// }
}

func switch_func() {
	// n := 1
	// switch n {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("I don't Know")
	// }
	
	/* switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't Know")
	} */

/* 	n := 6
	switch {
	case n > 0 && n < 4:
		fmt.Println("0 < n < 4")
	case n > 3 && n < 7:
		fmt.Println("3 < n < 7")
	} */

	// switch n := 2; n {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// case n > 3 && n < 6: // errorになる
	// 	fmt.Println("3 < n < 6")
	// default:
	// 	fmt.Println("I don't Know")
	// }
}

func anything(a interface{}) {
	// fmt.Println(a)

	switch v := a.(type) {
	case int:
		fmt.Println(v + 10000)
		fmt.Println("agument is Int")
	case string:
		fmt.Println(v + "!?")
		fmt.Println("agument is String")
	default:
		fmt.Println(v, "I don't Know")
	}
}
func type_assetion() {
	// anything("aaa")
	// anything(1)

	var x interface{} = 3
	i := x.(int)
	fmt.Println(i + 2)
	// fmt.Println(x + 2) // interface 型と int 型なので計算できない

	// f := x.(float64)
	// fmt.Println(f) // panic になり強制終了　

	f2, isFloat64 := x.(float64)
	fmt.Println(f2, isFloat64) // isFloat が false になっている

	// x = "string"  // "x is String"
	// x = [2]string{"1","2"} // "I don't Know"
	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
	} else if i, isString := x.(string); isString {
		fmt.Println(i, "x is String")
	} else {
		fmt.Println("I don't Know")
	}


	switch x.(type) {
	case int:
		fmt.Println("INT")
	case string:
		fmt.Println("STRING")
	default:
		fmt.Println("I don't Know")
	}

	switch v := x.(type) {
	case int:
		fmt.Println(v, "v is INT")
	case string:
		fmt.Println(v, "v is STRING")
	default:
		fmt.Println(v, "I don't Know")
	}

	anything("aaa") // aaa is String
	anything(1) // 1 is Int
}

func nested_for() {
// ラベル付きfor
/* Loop:
	for {
		for {
			for {
				fmt.Println("START")
				break Loop
			}
			fmt.Println("処理をしない")
		}
		fmt.Println("処理をしない")
	}
	fmt.Println("END") */
Loop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j > 1 {
				continue Loop
			}
			fmt.Println(i, j, i * j)
		}
		fmt.Println("処理をしない")
	}
}



func TestDefur() {
	defer fmt.Println("END") //関数が終了されるとき実行される
	fmt.Println("START")
}


func Rundefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func defer_function() {
	// defur 
	TestDefur()

	// defer func() {
	// 	fmt.Println("1")
	// 	fmt.Println("2")
	// 	fmt.Println("3")
	// }()

	Rundefer() // 3 2 1の順で表示される

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close() // 最後にファイルを閉じる

	file.Write([]byte("Hello"))
}

func panipani(){
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("runtime error")
	fmt.Println("start")
}

/* func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func go_routin() {
	// go goroutin
	go sub() // 並行処理できる
	go sub() // 並行処理できる

	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond)
	}
} */


func init() {
	fmt.Println("init")
}

func go_init(){
	fmt.Println("Main")
}