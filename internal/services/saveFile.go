package services

import "os"

func SaveToFile(filename string, content string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
