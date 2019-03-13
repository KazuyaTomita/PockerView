package main

import (
	"bufio"
	"fmt"
	"os"
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
		// TODO need to set inputScanner

	} else if config.Engines.Enable {
		// play games with multiple engines used
		fmt.Printf("multi-engine mode\n")
		panic("not implemented now. Can you send pull request?")
	} else {
		// CLI mode
		fmt.Printf("cli mode\n")
		inputScanner = bufio.NewScanner(os.Stdin)
	}

	ConnectEngine(inputScanner, config.Cli.Path)
}