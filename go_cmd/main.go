package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//os.Args

	var test int
	fmt.Println(test)

	var cmd = "ls"
	switch cmd {
	case "ls":
		fmt.Println("ls switch")
	case "ps":
		fmt.Println("ps switch")
		fallthrough
	case "ping":
		fmt.Println("ping switch")

	}

	//var breakCount = 6
	//for {
	//	if breakCount < 0 {
	//		break
	//	}
	//	fmt.Println(breakCount)
	//	time.Sleep(time.Second)
	//	breakCount--
	//}
	//fmt.Println("break count")
	//var count = 5
	//for count > 0 {
	//	fmt.Println(count)
	//	time.Sleep(time.Second)
	//	count--
	//}

	var num = 10
	var randNum int

	for {
		randNum = rand.Intn(100)
		if randNum > num {
			fmt.Println("is big", randNum)
		} else if randNum < num {
			fmt.Println("is small", randNum)
		} else {
			fmt.Println("is true", randNum)
			break
		}
	}

	var dayNum int
	dayNum = rand.Intn(10)
	switch dayNum {
	case 1:
		fmt.Println(dayNum)
	case 2, 3, 4, 5, 6, 8:
		fmt.Println("test", dayNum)
	default:
		fmt.Println("other", dayNum)

	}

}
