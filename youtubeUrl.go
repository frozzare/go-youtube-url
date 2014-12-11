package youtubeUrl

import (
	"regexp"
)

// Validate youtube url
func Valid(url string) bool {
	m, _ := regexp.MatchString(`((http://)?)(www\.)?((youtube\.com/)|(youtu.be)).+`, url)
	return m
}
