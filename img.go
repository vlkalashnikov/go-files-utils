package files

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func LoadImage(filename string) (img image.Image, fileType string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()

	return image.Decode(f)
}

func SaveImage(filename, fileType string, img image.Image) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	switch fileType {
	case "jpg", "jpeg":
		err = jpeg.Encode(f, img, &jpeg.Options{Quality: 100})
	default:
		err = png.Encode(f, img)
	}

	return err
}
