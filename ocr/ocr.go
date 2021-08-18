package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var pdfName string
var total int

func main() {

	fmt.Printf("Please enter pdf path: ")
	fmt.Scanln(&pdfName)
	fmt.Printf("Please enter pdf convert page: ")
	fmt.Scanln(&total)

	fileNameWithSuffix := path.Base(pdfName)
	fileType := path.Ext(fileNameWithSuffix)
	fileNameOnly := strings.TrimSuffix(fileNameWithSuffix, fileType)

	outputName := filepath.Join(".", "output", fileNameOnly)
	if err := ConvertPdfToJpg(pdfName, outputName, total); err != nil {
		log.Fatal(err)
	}

	access_token := Fetch_token()

	fd, err := os.ReadDir(filepath.Join(".", "output"))
	if err != nil {
		log.Fatal(err)
	}

	for _, ff := range fd { //遍历目录
		if ff.Type().IsRegular() && !strings.HasSuffix(ff.Name(), ".txt") {
			GeneralBasic(ff.Name(), access_token)
		}
	}
}
