package utils

import "os"

func SaveFile(fileName string, body []byte) error {
	return os.WriteFile(fileName, body, 0666)
}
