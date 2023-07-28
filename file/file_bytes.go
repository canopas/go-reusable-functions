package file

import (
	"bytes"
	"io"
	"os"
)

func GetFileBytes(fileName string) (*bytes.Buffer, error) {

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileBytes := bytes.NewBuffer(nil)
	_, err = io.Copy(fileBytes, file)

	return fileBytes, err
}
