package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		line := stdin.Text()
		fmt.Println(line)
	}
}

