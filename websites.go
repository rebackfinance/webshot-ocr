package webshot

import (
	"log"
	"time"
)

// Screenshot is the default screenshotter which can take the screenshot of webpages
func (p *Webshot) Screenshot(requestURL string, removeModals bool, sleepInterval time.Duration) ([]byte, error) {
	p.Webdriver.Get(requestURL)

	if removeModals {
		// Inject JavaScript to handle modals
		js := `
			// Function to remove modals
			function removeModals() {
				document.querySelectorAll('[role="dialog"], .modal, .popup, [data-modal]').forEach(el => el.remove());
			}
			
			// Remove existing modals
			removeModals();

			// Set up a MutationObserver to remove dynamically added modals
			const observer = new MutationObserver((mutations) => {
				removeModals();
			});
			observer.observe(document.body, { childList: true, subtree: true });

			console.log("Modal removal script injected.");
		`
		_, err := p.Webdriver.ExecuteScript(js, nil)
		if err != nil {
			log.Printf("Error injecting JavaScript to remove modals: %v", err)
		} else {
			log.Println("JavaScript to remove modals executed successfully.")
		}
	}

	time.Sleep(sleepInterval)
	imgByte, err := p.Webdriver.Screenshot()
	if err != nil {
		return nil, err
	}

	defer p.Service.Stop()
	defer p.Webdriver.Quit()

	return imgByte, nil
}
