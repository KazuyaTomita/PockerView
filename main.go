package main

import (
	"bufio"
	"fmt"
	"github.com/BurntSushi/toml"
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
type Config struct {
	Server ServerConfig `toml:"server"`
	Engines MultiEnginesConfig `toml:"multiEngines"`
	Cli CliConfig `toml:"cli"`
}

type ServerConfig struct {
	Enable bool `toml:"enable"`
	Ip  string `toml:"ip"`
	Port  string `toml:"port"`
	Path string `toml:"path"`
}

type MultiEnginesConfig struct {
	Enable bool `toml:"enable"`
	Number int `toml:"number"`
	Paths []string `toml:"paths"`
}

type CliConfig struct {
	Enable bool `toml:"enable"`
	Path string `toml:"path"`
}


func main() {
	
	var config Config
	// read config file
	_, parseErr := toml.DecodeFile("config.toml", &config)
	if parseErr != nil {
		panic(parseErr)
	}

	// check whether we can decide which mode is used
	if (config.Server.Enable && config.Engines.Enable) ||
		(config.Server.Enable && config.Cli.Enable) ||
		(config.Cli.Enable && config.Engines.Enable) {
		panic("not sure of which mode is used. please check config. Only single enable flag should be true.")
	}

	if config.Server.Enable {
		// communicate with TCP/IP server
		fmt.Printf("server mode\n")

	} else if config.Engines.Enable {
		// play games with multiple engines used
		fmt.Printf("multi-engine mode\n")

	} else {
		// CLI mode
		fmt.Printf("cli mode\n")
	}
	
	stdin := bufio.NewScanner(os.Stdin)

	// we prepare to execute an engine.
	cmd := exec.Command("go", "run", "dummy_engine.go")
	// pipe connecting to the engine's standard output
	stdoutPipe, stdoutErr := cmd.StdoutPipe()
	// pipe connecting to the engine's standard input
	stdinPipe, stdinErr := cmd.StdinPipe()

	if stdinErr != nil || stdoutErr != nil {
		panic("could not get stdinPipe or stdoutPipe")
	}

	scanner := bufio.NewScanner(stdoutPipe)
	var writer = bufio.NewWriter(stdinPipe)

	go writeInput(stdin, writer)
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