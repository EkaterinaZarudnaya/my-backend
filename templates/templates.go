package templates

import (
	_ "embed"
)

var (
	//go:embed index.html
	indexHtml string
	//go:embed upload.html
	uploadHtml string
)

func GetTemp() map[string]string {
	return map[string]string{
		"index":  indexHtml,
		"upload": uploadHtml,
	}
}
