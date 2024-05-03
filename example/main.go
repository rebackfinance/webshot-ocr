package main

import (
	"fmt"
	"os"

	webshot "github.com/rebackfinance/webshot-ocr"
)

func main() {
	config := webshot.NewConfig{
		Address:     "http://localhost",
		Port:        4444, // you can change accordingly to which ever port you wish
		BrowserName: webshot.FirefoxBrowser,
		DebugMode:   true,                  // set to true if you want to get the logs
		DriverPath:  "./creds/geckodriver", // your gekodriver path goes in here
	}

	driver, err := webshot.NewWebshot(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	url := "https://github.com/iqquee"

	byteImage, err := driver.Screenshot(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileName := "screenshot" + ".png"
	pngData, _ := os.Create(fileName)
	pngData.Write([]byte(byteImage))
}
