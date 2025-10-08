package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	runFor := true
	for runFor {
		blue := color.New(color.FgBlue)
		redERR := color.New(color.FgRed)
		mainTXT := "Type 'help' for a list of commands.\n" +
			"Hello, enter the command:"

		blue.Println(mainTXT)

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			redERR.Println(err)
			return
		}
		input = strings.TrimSpace(input)

		if input == "ext" {
			blue.Println("Goodbye!")
			runFor = false
		}

		switch input {
		case "help":
			helpTXT := "add - Adds a new task to the list; \n" +
				"show - Allows you to view the contents of a task file;\n" +
				"done - Marks a task as completed;\n" +
				"del - Delete task;\n" +
				"help - Shows brief information about the command;\n" +
				"ext - exit programm.\n"

			blue.Println(helpTXT)
		case "add":
			add()
		case "show":
			show()
		case "del":
			del()
		case "done":
			done()
		default:
			redERR.Println("unknown command")
		}
	}
}

//1428 ☆*: .｡. o(≧▽≦)o .｡.:*☆
