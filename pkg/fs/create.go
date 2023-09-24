package fs

import "os"

func CreateDir(path string) error {
	err := os.Mkdir(path, os.ModePerm)

	return err
}
