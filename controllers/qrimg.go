package controllers

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image"
	"image/png"
	"log"
	"os"
)

func WritePng(filename string, img image.Image) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	log.Println(file.Name())
}

func GetQrCode(str string) {
	code, err := qr.Encode(str, qr.L, qr.Unicode)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Encoded data:", code.Content())
	if str != code.Content() {
		log.Fatal("data differs")
	}
	code, err = barcode.Scale(code, 300, 300)
	if err != nil {
		log.Fatal(err)
	}
	WritePng("test.png", code)
}
