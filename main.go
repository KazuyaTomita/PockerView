package main

import (
	"bufio"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"os/exec"
)

type Config struct {
	Server ServerConfig
	Engines MultiEnginesConfig
	Cli CliConfig
}

type ServerConfig struct {
	Enable bool
	Ip  string
	Port  string
	Path string
}

type MultiEnginesConfig struct {
	Enable bool
	number int
	Path []string
}



type CliConfig struct {
	Enable bool
	Path string
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

	} else if config.Engines.Enable {
		// play games with multiple engines used

	} else {
		// CLI mode

	}




	fmt.Printf("Enable is :%t\n", config.Server.Enable)

	fmt.Printf("multiEngine Enable is :%t\n", config.Engines.Enable)
	fmt.Printf("Cli Enable is :%t\n", config.Cli.Enable)
	fmt.Printf("Port is :%s\n", config.Server.Port)

	fmt.Printf("%v\n", config.Server)
	fmt.Printf("%v\n", config.Engines)
	fmt.Printf("%v\n", config.Cli)



	stdin := bufio.NewScanner(os.Stdin)


	cmd := exec.Command("go", "run", "dummy_engine.go")
	stdoutPipe, stdoutErr := cmd.StdoutPipe()
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
// second one is used to write the input
func writeInput(scanner *bufio.Scanner, writer *bufio.Writer) {
	for scanner.Scan() {
		line := scanner.Text()
		writer.WriteString(line)
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