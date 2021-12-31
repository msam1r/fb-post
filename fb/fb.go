package fb

import (
	"log"
	"time"

	"github.com/msam1r/fb-post/code"
	"github.com/mxschmitt/playwright-go"
)

func Run(browser playwright.Browser, c code.Code, email, password string) {
	page, err := browser.NewPage()
	assertErrorToNilf("could not create page: %v", err)

	Login(page, email, password)
	_ = waitFor(page, 5)
	Confirm2FA(page, c)
	PublishPost(page)
}

func fillInput(page playwright.Page, id, name, value string) {
	entry, err := page.QuerySelector(id)
	if err != nil {
		log.Fatalf("could not get %s input. err: (%v)", name, err)
	}

	entry.Fill(value)
}

func assertErrorToNilf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func waitFor(page playwright.Page, seconds int) playwright.Response {
	var timeout float64 = float64(time.Second) * float64(seconds)
	resp, err := page.WaitForNavigation(playwright.PageWaitForNavigationOptions{
		Timeout: &timeout,
	})
	assertErrorToNilf("could not wait for navigation. err: (%v)", err)

	return resp
}
