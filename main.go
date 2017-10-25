package main

import (
	"log"
	"runtime"

	"github.com/sclevine/agouti"
)

func openCommand() string {
	command := ""
	os := runtime.GOOS

	if os == "linux" {
		command = "gnome-open"
	} else if os == "darwin" {
		command = "open"
	}

	return command
}

func main() {
	const pageImage = "/tmp/ame.png"

	driver := agouti.ChromeDriver()

	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	if err := page.Navigate("http://tokyo-ame.jwa.or.jp/"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
}
