package fb

import (
	"fmt"
	"log"

	"github.com/msam1r/fb-post/code"
	"github.com/mxschmitt/playwright-go"
)

func Confirm2FA(page playwright.Page, c code.Code) {
	page.Check("input") // wait till the page have an input element

	// Fill the confirmation input with code
	code := c.GetCode()
	fillInput(page, "#approvals_code", "approvals_code", code)

	// Click Submit code button
	entry, err := page.QuerySelector("#checkpointSubmitButton-actual-button")
	assertErrorToNilf("could not get button. err: (%v)", err)
	entry.Click()

	waitFor(page, 5)

saveBrowser:

	i := 1

	page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String(fmt.Sprintf("save-%d.png", i)),
	})

	if page.URL() == "https://m.facebook.com/login/checkpoint/" {
		submitAndContinue(page)
	}

	waitFor(page, 5)

	if page.URL() == "https://m.facebook.com/login/checkpoint/" {
		goto saveBrowser
	}

	log.Println("2FA done")
}

func submitAndContinue(page playwright.Page) {
	page.Check("button")
	// Click Continue (to save browser) & confirm that you are trying to login
	entry, err := page.QuerySelector("#checkpointSubmitButton-actual-button")
	assertErrorToNilf("could not get button. err: (%v)", err)
	entry.Click()
	fmt.Println("save browser done")
}
