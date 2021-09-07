package main

import (
	"fmt"
	"log"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()

	client.Languages = []string{"jpn"}
	client.SetImage("test/data/kanji/o.png")
	text, err := client.Text()
	if err != nil {
		log.Fatalln("文字を認識できませんでした。")
	}
	fmt.Println(text)
}
