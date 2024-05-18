package webshot

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/tebeka/selenium"
)

// install using sudo pacman -S xorg-server-xvfb

// NewWebshot sets up a new browser with the config provided
func NewWebshot(c NewConfig) (*Webshot, error) {
	opts := []selenium.ServiceOption{
		// selenium.StartFrameBuffer(),
		selenium.GeckoDriver(c.DriverPath),
		// firefox.
		selenium.Output(os.Stderr),
	}

	if c.DebugMode {
		selenium.SetDebug(true)
	}

	service, err := selenium.NewGeckoDriverService(c.DriverPath, c.Port, opts...)
	if err != nil {
		return nil, err
	}
	// defer service.Stop()

	var binaryLocation string
	if len(c.FirefoxBinary) == 0 {
		bl, err := getFirefoxBinaryLocation()
		if err != nil {
			return nil, err
		}
		binaryLocation = bl
	} else {
		binaryLocation = c.FirefoxBinary
	}

	// Connect to the WebDriver instance running locally
	caps := selenium.Capabilities{
		"browserName": c.BrowserName,
		"moz:firefoxOptions": map[string]interface{}{
			"args":   []string{"-headless"}, // Run in headless mode
			"binary": binaryLocation,        // Specify the path to the Firefox binary
		},
	}

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("%s:%d", c.Address, c.Port))
	if err != nil {
		return nil, err
	}

	fmt.Printf("The service ran successfully\n")
	return &Webshot{
		Webdriver: wd,
		Service:   service,
	}, nil
}

// Extend allow you to use use all of the functionalities provided by selenium
func (w *Webshot) Extend() selenium.WebDriver {
	return w.Webdriver
}

func getFirefoxBinaryLocation() (string, error) {
	var cmd *exec.Cmd

	// Determine the command to run based on the operating system
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("where", "firefox")
	default: // Linux, macOS, etc.
		cmd = exec.Command("which", "firefox")
	}

	// Execute the command
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// Convert the output to a string and trim any surrounding whitespace
	location := strings.TrimSpace(string(output))
	return location, nil
}
