package youtubeUrl

import (
	"regexp"
)

// Validate youtube url
func Valid(url string) bool {
	m, _ := regexp.MatchString(`((?:http://)?)(?:www\.)?(?:(youtube\.com/(\/watch\?(?:\=.*v=((\w|-){11}))|.+))|(youtu.be\/\w{11}))`, url)
	return m
}
