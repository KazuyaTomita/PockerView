package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("go", "run", "dummy_engine.go")
	stdout, stdoutErr := cmd.StdoutPipe()
	stdin, stdinErr := cmd.StdinPipe()

	if stdinErr != nil || stdoutErr != nil {
		panic("could not get stdin or stdout")
	}

	scanner := bufio.NewScanner(stdout)
	var writer = bufio.NewWriter(stdin)

	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			writer.WriteString("hoge\n")
			writer.Flush()
			time.Sleep(time.Second)
		}
	}()

	cmd.Run()
}