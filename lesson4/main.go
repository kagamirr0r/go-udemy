package main

import "fmt"

func main() {
	// int_type()
	// float_type()
	// bool_type()
	// string_type()
	// byte_type()
	// array_type()
	interface_type()
}

func int_type() {
	// 整数型
	var i int = 100
	fmt.Println(i)


	var i2 int64 = 200
	fmt.Println(i2)

	fmt.Println(i + 50)
	// fmt.Println(i + i2) // missmatched error

	fmt.Printf("%T\n", i2) // int64 型表示

	fmt.Printf("%T\n", int32(i2)) //int32 型キャスト

	fmt.Println(int64(i) + i2) // error にならない！
}

func float_type() {
	// 不動小数点型
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2 // float64になる
	fmt.Println(fl + fl64)

	fmt.Printf("%T\n", fl64) //float64
	fmt.Printf("%T\n", fl)	//float64

	var fl32 float32 = 1.2
	fmt.Println(fl32)	
	fmt.Printf("%T\n", fl32)	//float32

	fmt.Printf("%T\n", float64(fl32)) //float64

	zero := 0.0
	pinf := 1.0 / zero

	fmt.Println(pinf) //正の無限大

	ninf := -1.0 / zero
	fmt.Println(ninf) //負の無限大

	nan := zero / zero
	fmt.Println(nan) // nan
}

func bool_type() {
	var t, f bool = true, false
	fmt.Println(t, f)
}

func string_type() {
	var s string = "Hello Go"
	fmt.Println(s)
	fmt.Printf("%T\n", s) //string

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", s) //string

	fmt.Println(`test
	test
		test`)


	fmt.Println("\"")
	fmt.Println(`"`)

	fmt.Println(string(s[0])) //文字列はバイト
}

func byte_type() {
	//byte型
	byteA := []byte{72, 73}
	fmt.Println(byteA) //[72 73]

	fmt.Println(string(byteA)) //HI

	c := []byte("HI") 
	fmt.Println(c) //[72 73]
}

func array_type() {
	// 配列型
	var arr1 [3]int
	fmt.Println(arr1) // [0 0 0]

	fmt.Printf("%T\n", arr1) // [3]int

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2) //[A B ]

	arr3 := [3]int{1,2,3}
	fmt.Println(arr3) // [1 2 3]

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4) // [C D]
	fmt.Printf("%T\n", arr4) // [2]string ... は要素数をカウントしてくれる

	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])
	fmt.Println(arr2[2-1]) // B

	arr2[2] = "C"
	fmt.Println(arr2) // [A B C]

	// var arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5) //cannot use arr1 (variable of type [3]int) as type [4]int in assignment


	fmt.Println(len(arr1)) //3 配列の要素数
}


func interface_type() {
	// interface & nil

	var x interface{}
	fmt.Println(x) //<nil>

	// intefaze 型は互換性ありまくり
	x = 1
	fmt.Println(x) // 1
	x = 3.14
	fmt.Println(x) // 3.14
	x = "Go lang"
	fmt.Println(x) // Go lang
	x = [3]int{1,2,3}
	fmt.Println(x) // [1 2 3]

	// 計算はできない...
	// x = 2
	// fmt.Println(x + 3) // (mismatched types interface{} and int)
}
