package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		printMenuOptions()

		option := readOption()

		isRunning := handleOption(option)

		if !isRunning {
			// break
			os.Exit(0)
		}
	}
}

func handleOption(option int) bool {
	isRunning := true
	switch option {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Logs:")
	case 9:
		fmt.Println("Exiting...")
		isRunning = false
		break
	default:
		fmt.Println("Unknown option!")
	}
	return isRunning
}

func printMenuOptions() {
	fmt.Println("\n1 - Start Monitoring")
	fmt.Println("2 - Show logs")
	fmt.Print("9 - Exit application\n\n")
}

func readOption() int {
	fmt.Println("Choose an option to start:")
	var option int
	fmt.Scan(&option)
	return option
}
