package webshot

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/firefox"
)

// NewWebshot sets up a new browser with the config provided
func NewWebshot(c NewConfig) (*Webshot, error) {
	opts := []selenium.ServiceOption{
		// selenium.StartFrameBuffer(),
		selenium.GeckoDriver(c.DriverPath),
		selenium.Output(os.Stderr),
	}

	if c.DebugMode {
		selenium.SetDebug(true)
	}

	service, err := selenium.NewGeckoDriverService(c.DriverPath, c.Port, opts[0])
	if err != nil {
		return nil, err
	}

	var firefoxPath string
	// Check the operating system
	switch runtime.GOOS {
	case "windows":
		firefoxPath = findFirefoxWindows()
	case "darwin":
		firefoxPath = findFirefoxDarwin()
	case "linux":
		firefoxPath = findFirefoxLinux()
	default:
		fmt.Println("Unsupported operating system")
	}

	caps := selenium.Capabilities{"browserName": c.BrowserName}
	// Specify the path to the Firefox binary in capabilities
	fireFox := firefox.Capabilities{
		Binary: firefoxPath,
	}

	caps.AddFirefox(fireFox)

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

func findFirefoxWindows() string {
	// Look for Firefox in the default installation directory on Windows
	firefoxPath, err := exec.LookPath("firefox.exe")
	if err != nil {
		fmt.Println("Firefox not found on Windows:", err)
		return ""
	}
	return firefoxPath
}

func findFirefoxDarwin() string {
	// Look for Firefox in the default installation directory on macOS
	firefoxPath, err := exec.LookPath("/Applications/Firefox.app/Contents/MacOS/firefox")
	if err != nil {
		fmt.Println("Firefox not found on macOS:", err)
		return ""
	}

	return firefoxPath
}

func findFirefoxLinux() string {
	// Look for Firefox in the default installation directories on Linux
	firefoxPath, err := exec.LookPath("firefox")
	if err != nil {
		fmt.Println("Firefox not found on Linux:", err)
		return ""
	}
	return firefoxPath
}
