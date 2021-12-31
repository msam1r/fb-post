package main

import (
	"flag"

	"github.com/msam1r/fb-post/app"
	"github.com/msam1r/fb-post/code"
)

func main() {
	email := flag.String("email", "", "Email address")
	password := flag.String("pass", "", "Password")
	flag.Parse()

	c := code.Code{
		Config: &code.Config{
			CodesFile:     "txt/codes.txt",
			UsedCodesFile: "txt/used.txt",
		},
	}

	app.Run(c, *email, *password)
}
