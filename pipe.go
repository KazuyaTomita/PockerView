package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

// first argument is used to read input
// second one is used to write the input which is received through the scanner
func ConnectEngine(inputScanner *bufio.Scanner, path string) {
	// we prepare to execute an engine.
	cmd := exec.Command(path)
	// pipe connecting to the engine's standard output
	stdoutPipe, stdoutErr := cmd.StdoutPipe()
	// pipe connecting to the engine's standard input
	stdinPipe, stdinErr := cmd.StdinPipe()

	if stdinErr != nil || stdoutErr != nil {
		panic("could not get stdinPipe or stdoutPipe")
	}

	scanner := bufio.NewScanner(stdoutPipe)
	var writer = bufio.NewWriter(stdinPipe)

	go writeInput(inputScanner, writer)
	go printOutput(scanner)

	cmd.Run()
}

// first argument is used to read input
// second one is used to write the input which is received through the scanner
func writeInput(scanner *bufio.Scanner, writer *bufio.Writer) {
	for scanner.Scan() {
		line := scanner.Text()
		writer.WriteString(line)
		// newline code is really important when sending messages
		writer.WriteString("\n")
		writer.Flush()
	}
}

// receive output and print it
func printOutput(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
