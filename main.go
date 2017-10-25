package main

import (
	"log"
	"os/exec"
	"runtime"
	"time"

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
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	// Wait for display
	time.Sleep(2 * time.Millisecond)

	if err := page.Navigate("http://tokyo-ame.jwa.or.jp/"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}

	page.Screenshot(pageImage)

	exec.Command(openCommand(), pageImage).Start()
}
