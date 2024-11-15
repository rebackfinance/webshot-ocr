# webshot
webshot is a website screenshotting library in golang
### Screenshot taken using this library
<img src="./example/screenshot.png"></img>
# Get Started
In other to use this package, you need to first install `tesseract` on your machine and then download GeckoDriver for your os from [https://github.com/mozilla/geckodriver/releases](https://github.com/mozilla/geckodriver/releases).

### NOTE: The browser used in this package by default is firefox. Kindly install firefox if you don't have it on your machine already.
# Installation
This package can be installed by using the go command below.
```sh
go get github.com/rebackfinance/webshot-ocr@latest
```
# Quick start
```sh
# assume the following codes in example.go file
$ touch example.go
# open the just created example.go file in the text editor of your choice
```

# Screenshot
Screenshot is the default screenshotter which can take the screenshot of webpages

```go
package main

import (
	"fmt"
	"os"
	"time"

	webshot "github.com/rebackfinance/webshot-ocr"
)

func main() {
	config := webshot.NewConfig{
		Address:     "http://localhost",
		Port:        4444, // you can change accordingly to which ever port you wish
		BrowserName: webshot.FirefoxBrowser,
		DebugMode:   true, // set to true if you want to get the logs
		DriverPath:  "", // your gekodriver path goes in here
	}

	driver, err := webshot.NewWebshot(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	url := "https://google.com"
	sleepInverval := 4 * time.Second
	byteImage, err := driver.Screenshot(url, sleepInverval)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileName := "screenshot" + ".png"
	pngData, _ := os.Create(fileName)
	pngData.Write([]byte(byteImage))
}
```


# ImageProcessing
ImageProcessing does the optical character recognition(OCR)

This method processess an image and returns the text in that image in a .txt file.

```go
package main

import (
	"fmt"

	webshot "github.com/rebackfinance/webshot-ocr"
)

func main() {

    filePath := ""
    fileName := ""
	if err := webshot.ImageProcessing(filePath, fileName); err != nil {
		fmt.Println("Image processing err: \n", err)
		return
	}
}
```

# Extend
Extend allow you to use use all of the functionalities provided by selenium


You can use the Extend() to pretty much do whatever you wish as long as it's a functionality selenium supports. 
```go
package main

import (
	"fmt"

	webshot "github.com/rebackfinance/webshot-ocr"
)

func main() {
    config := webshot.NewConfig{
		Address:     "http://localhost",
		Port:        4444, // you can change accordingly to which ever port you wish
		BrowserName: webshot.FirefoxBrowser,
		DebugMode:   true, // set to true if you want to get the logs
		DriverPath:  "", // your gekodriver path goes in here
	}

	driver, err := webshot.NewWebshot(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	service := driver.Extend()
	imgBytes, err := service.Screenshot()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(imgBytes)
}
```