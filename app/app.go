package app

import (
	"log"

	"github.com/msam1r/fb-post/code"
	"github.com/msam1r/fb-post/fb"
	"github.com/mxschmitt/playwright-go"
)

func Run(c code.Code, email, password string) {
	pw, browser := initialize()

	fb.Run(browser, c, email, password)

	if err := browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}

	if err := pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}

func initialize() (*playwright.Playwright, playwright.Browser) {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
		return nil, nil
	}

	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
		return nil, nil
	}

	return pw, browser
}
