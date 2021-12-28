package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()

	client.Languages = []string{os.Getenv("LANG")}
	// TODO:DIテスト

	fileName := "./img/E054.png"
	if err := DownloadFile(fileName, "https://www.nichibenren.jp/member_general/font/E1BF.png"); err != nil {
		panic(err)
	}
	client.SetImage(fileName)
	text, err := client.Text()
	if err != nil {
		log.Fatalln("テキストに変換できませんでした。")
	}
	fmt.Println(text)

	if err := os.RemoveAll(fileName); err != nil {
		log.Fatalln(err)
	}
}

// DownloadFile ファイルダウンロード
func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func() {
		if err := out.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	_, err = io.Copy(out, resp.Body)

	return err
}
