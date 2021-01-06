package utils

import (
	"net/http"
	"os"
	"strings"
)

func GetFileType(fPath string) string {
	return strings.Split(GetFileContentType(fPath), "/")[0]
}

func GetFileContentType(fPath string) string {
	if f, err := os.Open(fPath); err == nil {
		buf := make([]byte, 512)
		if _, err = f.Read(buf); err == nil {
			return http.DetectContentType(buf)
		}
	}
	return ""
}