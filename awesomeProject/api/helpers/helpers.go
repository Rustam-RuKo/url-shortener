package helpers

import (
	"os"
	"strings"
)

func EnforceHTTP() string {
	if url[:4] != "http" {
		return "http:/" + url
	}
	return url
}

func RemoveDomainEditor(url string) bool {

	if url == os.Getenv("DOMAIN") {
		return False
	}

	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Replace(newURL, "/")[0]

	if newURL == os.Getenv("DOMAIN") {
		return False
	}
	return True
}
