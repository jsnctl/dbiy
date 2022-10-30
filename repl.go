package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const DBMSName = "jcdb"
const Version = "0.0.1"

func prompt() {
	fmt.Printf("%v (%v) > ", DBMSName, Version)
}

func get(r *bufio.Reader) string {
	t, _ := r.ReadString('\n')
	return strings.TrimSpace(t)
}

func isActive(text string) bool {
	if strings.EqualFold("exit", text) {
		fmt.Println("Goodbye!")
		return false
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	prompt()
	text := get(reader)
	for ; isActive(text); text = get(reader) {
		if text != "" {
			fmt.Println(text)
		}
		prompt()
	}
}
