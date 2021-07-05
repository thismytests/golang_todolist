package openLinksInBrowser

import (
	"fmt"
	"os"
)

func parser(fileName string) []string {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("err", err)
		os.Exit(-1)
	}

	defer file.Close()

	buffer := make([]byte, 1024*5)

	n, err := file.Read(buffer)

	if err != nil {
		fmt.Println("err reading", err)
	}

	result := string(buffer[:n])
	//fmt.Println("", string(buffer[:n]))

	arr := []string{}

	var item string

	for i := 0; i < len(result); i++ {
		currentItem := string(result[i])

		if currentItem == "\r" {
			arr = append(arr, item)
			item = ""
		} else {
			item = item + currentItem

		}

	}

	return arr
}
