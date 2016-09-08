package iox

import (
	"io"
	"os"
)

func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func CopyFile(srcFile, destFile string) error {
	file, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer file.Close()

	dest, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer dest.Close()

	io.Copy(dest, file)

	return nil
}
