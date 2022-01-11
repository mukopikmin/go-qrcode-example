package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("url")
	now := time.Now().UnixNano()
	filename := string(rune(now)) + ".png"

	qrc, err := qrcode.New(query)
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
		return
	}

	writer, err := standard.New(filename)
	if err != nil {
		fmt.Printf("standard.New failed: %v", err)
		return
	}

	if err = qrc.Save(writer); err != nil {
		fmt.Printf("could not save image: %v", err)
	}

	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(fileBytes)

	// err = os.Remove(filename)
	// if err != nil {
	// 	panic(err)
	// }
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
