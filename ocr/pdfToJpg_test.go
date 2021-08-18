package main

import (
	"log"
	"testing"
)

func TestConvertPdfToJpg(t *testing.T) {
	pdfName := "test.pdf"
	imageName := "test"

	if err := ConvertPdfToJpg(pdfName, imageName, 2); err != nil {
		log.Fatal(err)
	}
}
