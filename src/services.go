package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

var blueSrv = color.New(color.FgBlue)
var redERRsrv = color.New(color.BgRed)
var noCompletedTask string = "(not completed)"

func add() {
	var titleNameTask string
	blueSrv.Println("Enter name task file: ")
	fmt.Scan(&titleNameTask)

	readerSrv := bufio.NewReader(os.Stdin)
	readerSrv.ReadString('\n')

	filePath := "storage/" + titleNameTask + noCompletedTask + ".txt"
	file, err := os.Create(filePath)
	if err != nil {
		redERRsrv.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	blueSrv.Println("Enter the task contents:")
	for {
		inputSrv, err := readerSrv.ReadString('\n')
		if err != nil {
			redERRsrv.Println(err)
			break
		}
		inputSrv = strings.TrimSpace(inputSrv)
		if inputSrv == "" {
			break
		}

		_, err = file.WriteString(inputSrv + "\n")
		if err != nil {
			fmt.Println("Error writing file:", err)
			break
		}
	}
}

func show() {
	var showNameTask string
	blueSrv.Println("Enter name task file:")
	fmt.Scan(&showNameTask)

	filePath := "storage/" + showNameTask + noCompletedTask + ".txt"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		filePath = "storage/" + showNameTask + "(completed).txt"
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		redERRsrv.Println("Error read file:", err)
		return
	}

	blueSrv.Println(string(data))

}

func del() {
	var delNameTask string
	blueSrv.Println("Enter name task: ")
	fmt.Scan(&delNameTask)

	filePath := "storage/" + delNameTask + ".txt"

	err := os.Remove(filePath)
	if err != nil {
		redERRsrv.Println("Error delete file:", err)
	}
	fmt.Println("The file was successfully deleted.")
}

func done() {
	var doneNameTask string
	blueSrv.Println("Enter name task:")
	fmt.Scan(&doneNameTask)

	oldfilepath := "storage/" + doneNameTask + noCompletedTask + ".txt"
	newfilepath := "storage/" + doneNameTask + "(completed).txt"

	err := os.Rename(oldfilepath, newfilepath)
	if err != nil {
		redERRsrv.Println("Error rename file:", err)
		return
	}
	blueSrv.Println("Task marked as completed.")

}
