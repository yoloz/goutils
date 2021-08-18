package main

// //sudo apt install tesseract-ocr  libtesseract-dev

// import (
// 	"bufio"
// 	"fmt"
// 	"io/fs"
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"strings"

// 	"github.com/otiai10/gosseract/v2"
// )

// func GoTesseract(imageDir string) {
// 	client := gosseract.NewClient()
// 	defer client.Close()
// 	fd, err := os.ReadDir(filepath.Join(".", "output"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, ff := range fd { //遍历目录
// 		if ff.Type().IsRegular() && !strings.HasSuffix(ff.Name(), ".txt") {
// 			client.SetImage(ff.Name())
// 			text, err := client.Text()
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			writeFile(ff, text)
// 		}
// 	}

// }

// func writeFile(file fs.DirEntry, text string) {

// 	f, err := os.OpenFile(file.Name()+".txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	w := bufio.NewWriter(f)
// 	fmt.Fprintln(w, text)
// 	w.Flush()
// }
