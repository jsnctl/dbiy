package main

import (
	"bufio"
	"fmt"
	"os"
)

func prompt() {
	fmt.Print("dbiy (0.0.1) > ")
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		prompt()
		for reader.Scan() {
			text := reader.Text()
			fmt.Print(text)
			fmt.Println()
			prompt()
		}
	}
}
