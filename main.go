package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)


	cmd := exec.Command("go", "run", "dummy_engine.go")
	stdoutPipe, stdoutErr := cmd.StdoutPipe()
	stdinPipe, stdinErr := cmd.StdinPipe()

	if stdinErr != nil || stdoutErr != nil {
		panic("could not get stdinPipe or stdoutPipe")
	}

	scanner := bufio.NewScanner(stdoutPipe)
	var writer = bufio.NewWriter(stdinPipe)

	go printOutput(scanner)

	go writeInput(stdin, writer)

	cmd.Run()
}

// first argument is used to read input
// second one is used to write the input
func writeInput(scanner *bufio.Scanner, writer *bufio.Writer) {
	for scanner.Scan() {
		line := scanner.Text()
		writer.WriteString(line)
		writer.Flush()
		fmt.Println(line)
	}
}

// receive output and print it
func printOutput(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}