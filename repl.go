package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const DBMSName = "dbiy"
const Version = "0.0.1"

func prompt() {
	fmt.Printf("%v> ", DBMSName)
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
	fmt.Printf("%v version %v \n", DBMSName, Version)
}

func getAbout() {
	fmt.Printf("%v is a homemade DBMS \n", DBMSName)
}

type Record struct {
	ID string
	A  float64
	B  int
}

func doInsert(object string) {
	fmt.Printf("Insert (%v)\n", object)
}

func lookupCommand(text string) {
	commands := map[string]interface{}{
		"version": getVersion,
		"about":   getAbout,
		"insert":  doInsert,
	}
	command := strings.Fields(text)[0]
	if cmd, exists := commands[command]; exists {
		if command == "insert" {
			object := strings.TrimLeft(text, "insert ")
			doInsert(object)
		} else {
			cmd.(func())()
		}
	} else if text == "" {
		// nothing
	} else {
		fmt.Printf("Command %v not found \n", text)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	prompt()
	text := get(reader)
	for ; isActive(text); text = get(reader) {
		lookupCommand(text)
		prompt()
	}
}
