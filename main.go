package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	clearTerminal()

	printLogo()

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s?%s Project name: ", cyan, reset)

	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("erro:", err)
		return
	}

	name = strings.TrimSpace(name)

	project := FastForge{
		Name: name,
	}

	project.CreateProject()
}


func clearTerminal() {
	var cmd *exec.Cmd

	if os.Getenv("OS") == "Windows_NT" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("erro ao limpar terminal:", err)
	}
}

const (
	reset = "\033[0m"
	cyan  = "\033[36m"
	gray  = "\033[90m"
	green = "\033[32m"
)

func printLogo() {
	fmt.Printf("%s%s%s\n", cyan, `

  ______        _   ______
 |  ____|      | | |  ____|
 | |__ __ _ ___| |_| |__ ___  _ __ __ _  ___
 |  __/ _´ / __| __|  __/ _ \| ´__/ _´ |/ _ \
 | | | (_| \__ \ |_| | | (_) | | | (_| |  __/
 |_|  \__,_|___/\__|_|  \___/|_|  \__, |\___|
                                   __/ |
                                  |___/

`, reset)
}