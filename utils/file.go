package utils

import (
	"os"
	"path"
	"strings"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func GetFileNameWithoutExt(p string) string {
	if p == "" {
		return ""
	}
	_, filename := path.Split(p)
	ext := path.Ext(filename)
	return strings.TrimSuffix(filename, ext)
}
