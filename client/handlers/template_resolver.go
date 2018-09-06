package handlers

import (
	"os"
	"path/filepath"
)

func GetTemplatePath(rTempPath string) (s string) {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, rTempPath)
}
