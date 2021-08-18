package main

// sudo apt install libmagic-dev libmagickwand-dev

import (
	"strconv"

	"gopkg.in/gographics/imagick.v2/imagick"
)

// ConvertPdfToJpg will take a filename of a pdf file and convert the file into an
// image which will be saved back to the same location. It will save the image as a
// high resolution jpg file with minimal compression.
func ConvertPdfToJpg(pdfName string, imageName string, pdfSize int) error {

	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	// Must be *before* ReadImageFile,Make sure our image is high quality
	if err := mw.SetResolution(200, 200); err != nil {
		return err
	}

	// Load the image file into imagick
	if err := mw.ReadImage(pdfName); err != nil {
		return err
	}

	// Must be *after* ReadImageFile,Flatten image and remove alpha channel, to prevent alpha turning black in jpg
	if err := mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_FLATTEN); err != nil {
		return err
	}

	// Set any compression (100 = max quality)
	if err := mw.SetCompressionQuality(90); err != nil {
		return err
	}
	// Convert into JPG
	if err := mw.SetFormat("jpg"); err != nil {
		return err
	}
	for i := 0; i < pdfSize; i++ {
		mw.SetIteratorIndex(i)
		// Save File
		if err := mw.WriteImage(imageName + "_" + strconv.Itoa(i) + ".jpg"); err != nil {
			return err
		}
	}
	return nil
}
