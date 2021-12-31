package fb

import (
	"log"

	"github.com/mxschmitt/playwright-go"
)

func PublishPost(page playwright.Page) {
	page.Goto("https://m.facebook.com/home.php")
	page.Check("div")

	// Click on the text area to open the post popup
	entry, err := page.QuerySelector("div._4g34._6ber._78cq._7cdk._5i2i._52we")
	if err != nil {
		log.Fatalf("Error while trying to open the post window. err: (%v)", err)
	}
	entry.Click()

	// wait for navigation with timeout 10 seconds
	waitFor(page, 10)
	// wait till the page have textarea input
	page.Check("textarea")

	path := "/home/mohamed/go/src/github.com/msam1r/fb-post/submit.js"
	entry, err = page.AddScriptTag(playwright.PageAddScriptTagOptions{
		Path: &path,
	})
	if err != nil {
		log.Fatalf("Error while creating & running the script. err: (%v)", err)
	}
	_ = entry

	page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("published.png"),
	})

	log.Println("Post published")
}
