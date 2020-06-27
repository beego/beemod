package jwt

import (
	"os"
)

func FileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}
