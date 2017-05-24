package controllers

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image"
	"image/png"
	"log"
	"os"
)

func WritePng(filename string, img image.Image) string {
	file, err := os.Create("./static/images/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	log.Println(file.Name())
	return "/static/images/" + filename
}

func GetQrCode(str string) string {
	code, err := qr.Encode(str, qr.L, qr.Unicode)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Encoded data:", code.Content())
	if str != code.Content() {
		log.Fatal("data differs")
	}
	code, err = barcode.Scale(code, 200, 200)
	if err != nil {
		log.Fatal(err)
	}
	return WritePng("test.png", code)
}
