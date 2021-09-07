package main

import (
	"fmt"
	"log"
	"os"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()

	client.Languages = []string{os.Getenv("LANG")}
	client.SetImage("test/data/kanji/o.png")
	text, err := client.Text()
	if err != nil {
		log.Fatalln("テキストに変換できませんでした。")
	}
	fmt.Println(text)
}
