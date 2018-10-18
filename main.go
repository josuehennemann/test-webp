package main

import (
	"bytes"
	"fmt"
	"github.com/chai2010/webp"
	"image/jpeg"
	"os"
)

var invalidArgs = fmt.Errorf("Invalid args")

func main() {

	if len(os.Args) != 2 {
		checkError(invalidArgs)
	}

	fileName := os.Args[1]

	//Read image
	file, er := os.Open(fileName)
	checkError(er)
	fmt.Println("Reading jpeg....")

	jpegImage, er := jpeg.Decode(file)
	checkError(er)

	fmt.Println("Converting...")

	var b bytes.Buffer
	err := webp.Encode(&b, jpegImage, nil)
	checkError(err)

	//create and save file webp
	f, er := os.Create(fileName + ".webp")
	checkError(er)
	f.Write(b.Bytes())
	f.Close()

	fmt.Println("Finish")
}

func checkError(e error) {
	if e == nil {
		return
	}
	fmt.Println(e)
	os.Exit(1)
}
