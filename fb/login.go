package fb

import (
	"log"

	"github.com/mxschmitt/playwright-go"
)

func Login(page playwright.Page, email, password string) {
	_, err := page.Goto("https://m.facebook.com/")
	assertErrorToNilf("could not goto: %v", err)

	// Fill email
	fillInput(page, "#m_login_email", "email", email)

	// Fill password
	fillInput(page, "#m_login_password", "password", password)

	// Click Login u_0_d_Kt
	entry, err := page.QuerySelector("button[name=login]")
	assertErrorToNilf("could not get button. err: (%v)", err)
	entry.Click()

	log.Println("Login Done")
}
