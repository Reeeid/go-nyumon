package main

import (
	"fmt"
	"os"
)

func main() {
	sampleFile := "./sample.txt"
	CheckFileExist(sampleFile)

	notExistFile := "./abc.txt"
	CheckFileExist(notExistFile)
}

func CheckFileExist(path string) {
	_, err := os.Stat(path)
	if err == nil {
		fmt.Print("ファイルが見つかりました")
	} else {
		fmt.Print("ファイルが見つかりませんでした")
	}
}
