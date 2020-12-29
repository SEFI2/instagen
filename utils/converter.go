package utils

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
)

func PNGtoJPG(pngFilename, jpgFilename string) error {
	// Decode
	pngFile, err := os.Open(pngFilename)
	if err != nil {
		fmt.Println("Cannot open the file. Error: ", err)
		return err
	}
	defer pngFile.Close()
	image, err := png.Decode(pngFile)
	if err != nil {
		return err
	}

	// Create file
	jpgFile, err := os.Create(jpgFilename)
	if err != nil {
		return err
	}
	defer jpgFile.Close()

	// Encode
	var opt jpeg.Options
	opt.Quality = 80
	if err := jpeg.Encode(jpgFile, image, &opt); err != nil {
		return err
	}
	return nil
}
