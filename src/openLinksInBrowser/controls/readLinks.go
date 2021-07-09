package controls

import (
	"fmt"
	"os"
)

func Parser(fileName string, separator string) ([]string, error) {
	file, err := os.Open(fileName)

	defer file.Close()

	if err != nil {
		fmt.Println("err", err)
		os.Exit(-1)
	}

	buffer := make([]byte, 1024*5)

	n, err := file.Read(buffer)

	if err != nil {
		fmt.Println("err reading", err)
	}

	result := string(buffer[:n])

	arr, err := convertStringToArr(result, separator)

	if err != nil {
		fmt.Println("err", err)
	}

	return arr, nil
}

func convertStringToArr(str string, separator string) ([]string, error) {
	var arr []string

	var item string

	for i := 0; i < len(str); i++ {
		currentItem := string(str[i])

		if currentItem == separator {
			arr = append(arr, item)
			item = ""
		} else {
			item = item + currentItem

		}

	}

	return arr, nil
}
