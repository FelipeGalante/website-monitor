package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for {
		printMenuOptions()

		option := readOption()

		handleOption(option)
	}
}

func handleOption(option int) {
	switch option {
	case 1:
		startMonitoring()
	case 2:
		fmt.Println("Logs:")
	case 9:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Unknown option!")
	}
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

func startMonitoring() {
	fmt.Println("Monitoring...\n")
	URLs := []string{
		"https://random-status-code.herokuapp.com",
		"https://nemo.levven.com/healthcheck",
	}

	for _, url := range URLs {
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error on request: ...", err)
		}
		fmt.Println(resp.Status, "\n")
	}
}
