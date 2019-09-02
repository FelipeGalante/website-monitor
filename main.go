package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const usage = `
Options:
  -h --help        Show this screen.
  -v --version     Show version.
`

const configFile = "config.json"

type URLs struct {
	URLs []string `json:"urls"`
}

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

	const (
		layoutISO = "2006-01-02"
		layoutUS  = "January 2, 2006 3:04 PM"
	)

	t := time.Now()

	fmt.Println("\n", t.Format(layoutUS))
	fmt.Println("----------MENU----------")
	fmt.Println("1 - Start Monitoring")
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

	URLs := createListURL()

	for i := 0; i < 5; i++ {
		executeMonitoring(URLs)
		time.Sleep(5 * time.Second)
	}

}

func executeMonitoring(urls URLs) {
	for i := 0; i < len(urls.URLs); i++ {
		url := urls.URLs[i]
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error on request: ...", err)
		}
		fmt.Println(resp.Status, "\n")
	}
}

func readConfigFile() []byte {
	jsonFile, err := os.Open(configFile)

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}

func createListURL() URLs {
	byteValue := readConfigFile()
	var urls URLs

	json.Unmarshal(byteValue, &urls)

	return urls
}
