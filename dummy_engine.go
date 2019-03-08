package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		line := stdin.Text()
		fmt.Println(line)
		fmt.Println("I am dummy engine")
	}
	//for stdin.Scan() {
	//	fmt.Println("dymmy_engine")
	//	switch stdin.Text() {
	//	case "is_ready":
	//		fmt.Println("hoge")
	//	default:
	//		fmt.Println("other")
	//	}
	//}
}

