package main

import (
	"fmt"
	"time"
)

// スライス
// 宣言、操作
func main() {
	var sl []int
	fmt.Println(sl)

	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)

	sl3 := []string{"apple", "banana"}
	fmt.Println(sl3)

	sl4 := make([]int, 5)
	fmt.Println(sl4)

	sl2[0] = 1000
	fmt.Println(sl2) // [1000 200]

	sl5 := []int{1,2,3,4,5}
	fmt.Println(sl5)

	fmt.Println(sl5[0]) // 1
	fmt.Println(sl5[1:4]) // [2 3 4] // 範囲指定
	fmt.Println(sl5[:3]) // [1 2 3] // 〜まで
	fmt.Println(sl5[2:]) // [3 4 5] // から〜
	fmt.Println(sl5[:]) // [1 2 3 4 5] 全て
	fmt.Println(sl5[len(sl5) - 1]) // 5 最後の値
	fmt.Println(sl5[0 : len(sl5) - 1]) // [1 2 3 4] 最初から最後の値を抜いたスライス

	slice_append()
	slice_copy()
	slice_for()
	variadic_arguments()

	// map
	map_main()
	map_for()


	//channel
	channel_main()
	channel_go_routin()
	channel_close()
	channel_for()
	channel_select()
}


func slice_append() {
	sl := []int{100, 200}
	fmt.Println(sl)

	sl = append(sl, 300)
	fmt.Println(sl)

	sl = append(sl, 400, 500, 600)
	fmt.Println(sl)

	sl2 := make([]int, 5)
	fmt.Println(sl2)

	fmt.Println(len(sl2))
	fmt.Println(cap(sl2)) // 容量

	sl3 := make([]int, 5, 10)
	fmt.Println(cap(sl3))

	sl3 = append(sl3, 1,2,3,4,5,6,7)
	fmt.Println(cap(sl3)) // 20になってしまう！！
}


func slice_copy() {
	sl := []int{100, 200}
	sl2 := sl

	sl2[0] = 1000
	fmt.Println(sl) //[1000 200] 元のslが書きかわっている！！

	var i int = 10
	i2 := i
	i2 = 1000

	fmt.Println(i, i2) //10 1000

	sl3 := []int{1,2,3,4,5}
	sl4 := make([]int, 5, 10)
	fmt.Println(sl4) // [0 0 0 0 0]

	n := copy(sl4, sl3)
	fmt.Println(sl4) // [1 2 3 4 5]
	fmt.Println(n) //copyに成功した数
}

func slice_for() {
	// for
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	for _,v := range sl {
		fmt.Println(v)
	}

	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}
}

func Sum(s ...int) int {
	n := 0

	// sはスライスになる
	for _, v := range s { 
		n += v
	}

	return n
}

func variadic_arguments() {
	// 可変長引数

	fmt.Println(Sum(1,2,3,4,5)) // 15
	fmt.Println(Sum(1,2,3,4,5,6,7,8,9,10)) // 55

	sl := []int{10,11,12}
	fmt.Println(Sum(sl...)) // 33 スライスも渡せる
}

func map_main() {
	// マップ
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	m2 := map[string]int{"A": 1000, "B": 2000}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)


	m4 := make(map[int]string)
	fmt.Println(m4)
	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4)

	fmt.Println(m["A"]) //100
	fmt.Println(m4[2]) //USA
	fmt.Println(m4[3]) //  （空文字が帰ってくる）

	s1, ok1 := m4[1]
	fmt.Println(s1, ok1) //JAPAN true (okには所得できたかどうかが入っている)

	s2, ok2 := m4[2]
	fmt.Println(s2, ok2)

	s3, ok3 := m4[3]
	if !ok3 {
		fmt.Println("Fuck my enemies")
	}
	fmt.Println(s3, ok3) // false

	m4[2] = "US"
	fmt.Println(m4[2]) 

	m4[3] = "CHINA"
	fmt.Println(m4) 

	delete(m4, 3)
	fmt.Println(m4)

	fmt.Println(len(m4))
}

func map_for() {
	m := map[string]int {
		"Apple": 100,
		"Banana": 150,
	}

	for k,v := range m {
		fmt.Println(k,v)
	}
}


// channel
func channel_main() {
	var ch1 chan int

	// //受信専用
	// var ch2 <-chan int

	// //送信専用
	// var ch3 chan<- int

	ch1 = make(chan int)
	ch2 := make(chan int)

	fmt.Println(cap(ch1)) // 0
	fmt.Println(cap(ch2)) // 0

	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3)) // 5

	ch3 <- 1
	fmt.Println(len(ch3)) // 1

	ch3 <- 2
	ch3 <- 3
	fmt.Println("length", len(ch3)) // 3

	i := <- ch3
	fmt.Println(i) // 1　　受け取れてる（先入れ先だし）
	fmt.Println("length", len(ch3)) //2 減っとる！！

	i2 := <- ch3
	fmt.Println(i2) // 2　　受け取れてる（先入れ先だし）
	fmt.Println("length", len(ch3))  //1 また減っとる！！
}


func reciver(c chan int) {
	for {
		i := <- c
		fmt.Println(i)
	}
}
func channel_go_routin() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// fmt.Println(<-ch) // deadlock

	go reciver(ch1)
	go reciver(ch2)

	i := 0
	for i < 100 {
		ch1 <- i
		ch2 <- i

		time.Sleep(50 * time.Millisecond)
		i++
	}
}


func receiver2(name string, c chan int) {
	for {
		i, ok := <-c

		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func channel_close() {
	ch1 := make(chan int, 2)

	close(ch1)

	// ch1 <- 1 //panic: send on closed channel
	// fmt.Println(<- ch1) // 0 受信はできる
	i, ok := <-ch1
	fmt.Println(i,ok) // 0 false falseはchannelがcloseされているということ

	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	fmt.Println(<-ch2)

	i2, ok2 := <-ch2
	fmt.Println(i2, ok2) // 2 true

	ch3 := make(chan int, 2)
	go receiver2("1.goroutin", ch3)
	go receiver2("2.goroutin", ch3)
	go receiver2("3.goroutin", ch3)

	i3 := 0
	for i3 < 100 {
		ch3 <-1
		i3++
	}
	close(ch3)
	time.Sleep(3 * time.Second)
}

func channel_for() {
	// channel
	// for

	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1) //deadlockを防ぐため
	for i := range ch1 {
		fmt.Println(i)
	}
}

func channel_select() {
	//channel
	//select

	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	// ch2 <- "A"
	// ch1 <- 1
	// ch2 <- "B"
	// ch1 <- 2

	// v1 := <- ch1
	// v2 := <- ch2
	// fmt.Println(v1) // 値がないのでdeadlock
	// fmt.Println(v2) // 実行されない

	select {
	case v1 := <- ch1:
		fmt.Println(v1 + 1000)
	case v2 := <- ch2:
		fmt.Println(v2 + "!?")
	default:
		fmt.Println("どちらでもない")
	}

	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	go func() {
		for {
			i := <- ch3
			ch4 <- i * 2
		}
	}()

	go func() {
		for {
			i2 := <- ch4
			ch5 <- i2 - 1
		}
	}()

	n := 0

	for {
		select {
		case ch3 <- n:
			n++
		case i3 := <-ch5:
			fmt.Println("received", i3)
		}
		if n > 100 {
			break
		}
	}

}

