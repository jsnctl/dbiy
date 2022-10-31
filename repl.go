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
	fmt.Printf("%v > ", DBMSName)
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

func getVersion() {
	fmt.Printf("jcdb version %v \n", Version)
}

func getAbout() {
	fmt.Printf("jcdb is a homemade DBMS \n")
}

func main() {
	commands := map[string]interface{}{
		"version": getVersion,
		"about":   getAbout,
	}
	reader := bufio.NewReader(os.Stdin)
	prompt()
	text := get(reader)
	for ; isActive(text); text = get(reader) {
		if text == "" {
		} else if command, exists := commands[text]; exists {
			command.(func())()
		} else {
			fmt.Printf("Command %v not found \n", text)
		}
		prompt()
	}
}
