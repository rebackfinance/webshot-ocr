package main

import (
	"fmt"
	"os"
	"time"

	webshot "github.com/rebackfinance/webshot-ocr"
)

func main() {
	fmt.Println("Starting...")
	config := webshot.NewConfig{
		Address:     "http://localhost",
		Port:        4447, // you can change accordingly to which ever port you wish
		BrowserName: webshot.FirefoxBrowser,
		DebugMode:   true,                  // set to true if you want to get the logs
		DriverPath:  "./creds/geckodriver", // your gekodriver path goes in here
	}

	fmt.Println("Proceeding with setting up driver...")
	driver, err := webshot.NewWebshot(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Preparing to visit website...")
	url := "https://www.instagram.com/p/C7M9ciTsuS8/"
	sleepInverval := 10 * time.Second
	byteImage, err := driver.Screenshot(url, true, sleepInverval)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success...")

	fileName := "https://www.instagram.com/p/C7M9ciTsuS8" + ".png"
	pngData, _ := os.Create(fileName)
	pngData.Write([]byte(byteImage))
}
