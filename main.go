package main

import (
	"fmt"
)

func main() {

	for {
		isRunning := true
		fmt.Println("\n1 - Start Monitoring")
		fmt.Println("2 - Show logs")
		fmt.Print("9 - Exit application\n\n")

		var option int
		fmt.Scan(&option)

		switch option {
		case 1:
			fmt.Println("Monitoring...")
		case 2:
			fmt.Println("Monitoring...")
		case 9:
			fmt.Println("Exiting...")
			isRunning = false
			break
		default:
			fmt.Println("Unknown option!")
		}
		if !isRunning {
			break
		}
	}

}
