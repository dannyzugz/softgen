package services

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func CompressFolder(folderPath string, zipWriter *zip.Writer) error {

	// Loop through all files and subfolders in the folder
	err := filepath.Walk(folderPath, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Create a new entry in the ZIP file for each file or folder
		header, err := zip.FileInfoHeader(fileInfo)
		if err != nil {
			return err
		}
		header.Name = filePath
		if fileInfo.IsDir() {
			header.Name += "/"
		}

		// Write the entry to the ZIP file
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// If the input is a file, copy its contents to the ZIP file
		if !fileInfo.IsDir() {
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
