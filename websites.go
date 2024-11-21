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
			(function () {
				// Function to remove modals and overlays
				function removeModalsAndOverlays() {
					const selectors = [
						'[role="dialog"]', '.modal', '.popup', '[data-modal]',
						'[class*="overlay"]', '[class*="backdrop"]', 
						'[id*="modal"]', '[id*="overlay"]', 
						'[aria-modal="true"]', '[role="presentation"]',
						'[data-testid="dialog"]', '[aria-label="Close"]'
					];

					// Remove elements matching the selectors
					selectors.forEach((selector) => {
						document.querySelectorAll(selector).forEach((el) => {
							el.remove();
						});
					});

					// Hide stubborn overlays based on high z-index
					document.querySelectorAll('*').forEach((el) => {
						const style = window.getComputedStyle(el);
						if ((style.position === 'fixed' || style.position === 'absolute') && parseInt(style.zIndex) > 999) {
							el.style.display = 'none';
						}
					});
				}

				// Execute the function immediately
				removeModalsAndOverlays();

				// Set up a MutationObserver to handle dynamically added overlays
				const observer = new MutationObserver(removeModalsAndOverlays);
				observer.observe(document.body, { childList: true, subtree: true });

				console.log("Overlay removal script injected and running.");
			})();
		`

		_, err := p.Webdriver.ExecuteScript(js, nil)
		if err != nil {
			log.Printf("Error injecting JavaScript to remove modals and overlays: %v", err)
		} else {
			log.Println("JavaScript to remove modals and overlays executed successfully.")
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
