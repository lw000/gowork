// qrcode_test project main.go
package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	qrcode "github.com/skip2/go-qrcode"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"

	ttqrcode "github.com/tuotoo/qrcode"
)

func writePng(filename string, img image.Image) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}
}

func createQR(tx string) {
	str := tx
	code, err := qr.Encode(str, qr.L, qr.Unicode)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Encoded data:", code.Content())

	if str != code.Content() {
		log.Fatal("data differs")
	}

	code, err = barcode.Scale(code, 300, 300)
	if err != nil {
		log.Fatal(err)
	}

	writePng("test.png", code)
}

func createQR1(tx string) {
	err := qrcode.WriteFile(tx, qrcode.Medium, 256, "test1.png")
	if err != nil {
		log.Fatal(err)
	}
}

func readQR(qrfile string) {
	fi, err := os.Open(qrfile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fi.Close()

	qrmatrix, err := ttqrcode.Decode(fi)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(qrmatrix.Content)
}

func main() {
	createQR("www.baidu.com")
	createQR1("www.google.com")

	readQR("test.png")
	readQR("test1.png")
}
