package util

import (
	"os"
	"path/filepath"
)

func GetPort() string {
	exePath, err := os.Executable()
	PanicIfErr(err)

	dirPath := filepath.Dir(exePath)
	path := filepath.Join(dirPath, "port.conf")

	content, err := os.ReadFile(path)
	PanicIfErr(err)

	return string(content)
}
