package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	//
	var a int
	fmt.Scan(&a)
	fmt.Printf("%d\n",a )
	//

	cmd := exec.Command("wc")
	stdin, _ := cmd.StdinPipe()
	fmt.Printf("readしたよ")
	io.WriteString(stdin, "hogehoge\\0")
	stdin.Close()
	fmt.Printf("closeしなかったよ")
	//?out, _ := cmd.Output()
	// fmt.Printf("結果: %s", out)
	fmt.Printf("clofafafaf")

}

