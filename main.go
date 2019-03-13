package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

/*
We design a poker AI by using two components. Generally speaking, The first one is a GUI program and the second one is a AI engine.
By doing so, we can separate interface parts and actual search or estimation parts. And this program is a GUI program.
However, implementing GUI by Golang is not easy now. So, this program is just cli tool.
We want to offer three types of interface.

(i) CLI. In the cli, a user just types commands such as "is_ready" and then this CLI program passes it to an engine.
After that, the engine executes something and sends messages. The CLI program shows the messages.

(ii) multi-engine case. Sometimes we want to make an engine battle other engine to evaluate their ratings. So, we support a multi-engine game.
In the case, a user needs to specify how many engines are used and which engines are used.

(iii) Communication with TCP/IP server which implements a certain protocol. Most official games are this type.
Under the hood, we create a socket and need to use some famous system-call.

 */


func main() {
	
	var config Config
	ReadConfig(&config)

	var inputScanner *bufio.Scanner

	if config.Server.Enable {
		// communicate with TCP/IP server
		fmt.Printf("server mode\n")

	} else if config.Engines.Enable {
		// play games with multiple engines used
		fmt.Printf("multi-engine mode\n")

	} else {
		// CLI mode
		fmt.Printf("cli mode\n")
		inputScanner = bufio.NewScanner(os.Stdin)
	}

	// we prepare to execute an engine.
	cmd := exec.Command(config.Cli.Path)
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