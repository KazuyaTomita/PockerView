package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	cmd := exec.Command("sleep 7 && echo 'hoge'")
	stdin, stdinErr := cmd.StdinPipe()
	stdout, stdoutErr := cmd.StdoutPipe()
	if stdinErr != nil || stdoutErr != nil {
		panic("could not get stdin or stdout")
	}
	fmt.Printf("write something\n" )
	io.WriteString(stdin, "hogeaaaa ddd")
	buffer := make([]byte,1024)
	fmt.Printf("read buffer\n" )
	// text ,_ := stdout.Read(buffer)
	io.ReadFull(stdout, buffer)
	fmt.Printf("print text from the command\n" )
	fmt.Printf("OUTPUT=%s", buffer[3])
	//
	//var a int
	//fmt.Scan(&a)
	//fmt.Printf("%d\n",a )
	//exec.Command("ls", "-la").Run()
	//
}

