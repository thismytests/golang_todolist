package main

import (
	"fmt"
	"log"
	"openLinksInBrowser/model"
	"os/exec"
	"runtime"
)

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	separator := "\n"

	links, err := model.Parser("./links.txt", separator)
	if err != nil {
		fmt.Println("err", err)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println("links[i]", links[i])
		openbrowser(links[i])
	}

}
