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
			// Function to remove modals and overlays
			function removeModalsAndOverlays() {
				// Remove modal elements
				document.querySelectorAll('[role="dialog"], .modal, .popup, [data-modal]').forEach(el => el.remove());

				// Remove known overlay elements
				document.querySelectorAll('.overlay, .backdrop, [class*="overlay"], [class*="backdrop"]').forEach(el => el.remove());

				// Specific handling for Instagram and Facebook overlays
				document.querySelectorAll('[class*="root"], [class*="layer"]').forEach(el => {
					const style = window.getComputedStyle(el);
					if (style.position === 'fixed' || style.zIndex > 1000) {
						el.remove(); // Remove overlay-like elements
					}
				});

				// General fallback: Hide stubborn overlays with high z-index
				document.querySelectorAll('*').forEach(el => {
					const style = window.getComputedStyle(el);
					if (style.position === 'fixed' && style.zIndex > 1000) {
						el.style.display = 'none'; // Force-hide high z-index elements
					}
				});
			}

			// Execute removal on initial load
			removeModalsAndOverlays();

			// Set up a MutationObserver to handle dynamically added modals and overlays
			const observer = new MutationObserver(() => {
				removeModalsAndOverlays();
			});

			observer.observe(document.body, { childList: true, subtree: true });

			console.log("Modal and overlay removal script injected.");
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
