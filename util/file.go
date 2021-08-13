package util

import (
	"io/ioutil"
	"net/http"
	"os"
)

func GetFileContentType(f *os.File) (string, error) {
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return http.DetectContentType(data), nil
}
