package storage

import (
	"os"
)

func GetFileData(fileName string) ([]byte, error) {
	if err := createIfNotCreated(fileName); err != nil {
		return []byte{}, err
	}

	fileData, err := os.ReadFile(fileName)
	if err != nil {
		return []byte{}, err
	}

	return fileData, nil
}

func WriteFileData(fileName string, data []byte) error {
	if err := createIfNotCreated(fileName); err != nil {
		return err
	}

	if err := os.WriteFile(fileName, data, 0755); err != nil {
		return err
	}

	return nil
}

func AppendFileData(fileName string, data []byte) error {
	if err := createIfNotCreated(fileName); err != nil {
		return err
	}

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		file.Close()
		return err
	}
	defer file.Close()

	if _, err := file.Write(data); err != nil {
		return err
	}

	return nil
}

func createIfNotCreated(fileName string) error {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	return nil
}
