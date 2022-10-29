package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func prompt() {
	fmt.Print("dbiy (0.0.1) > ")
}

func get(r *bufio.Reader) string {
	t, _ := r.ReadString('\n')
	return strings.TrimSpace(t)
}

func shouldContinue(text string) bool {
	if strings.EqualFold("exit", text) {
		return false
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	prompt()
	text := get(reader)
	for ; shouldContinue(text); text = get(reader) {
		fmt.Println(text)
		prompt()
	}
}
