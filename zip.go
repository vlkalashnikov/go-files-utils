package files

import (
	"archive/zip"
	"io"
	"os"
)

/*Создает архив на основе указанного файла*/
func CreateZipFile(fromfile, toFile string) error {
	archive, err := os.Create(toFile)

	if err != nil {
		return err
	}
	defer archive.Close()

	zipWriter := zip.NewWriter(archive)

	fileToZip, err := os.Open(fromfile)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
