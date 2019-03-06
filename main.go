package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func readStuff(scanner *bufio.Scanner) {
	for scanner.Scan() {
		fmt.Println("Performed Scan")
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func main() {
	cmd := exec.Command("wc")
	stdin, stdinErr := cmd.StdinPipe()
	stdout, stdoutErr := cmd.StdoutPipe()
	if stdinErr != nil || stdoutErr != nil {
		panic("could not get stdin or stdout")
	}
	io.WriteString(stdin, "hogeaaaa ddd")
	stdin.Close()
	err := cmd.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start err=%v", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(stdout)
	fmt.Println("Scanner created")

	defer cmd.Wait()

	go readStuff(scanner)
}