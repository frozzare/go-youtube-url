package youtubeUrl

import (
	"testing"

	"github.com/frozzare/go-assert"
)

func TestYoutubeUrls(t *testing.T) {
	assert.Equal(t, true, Valid("https://www.youtube.com/watch?v=Xq7z6WpeB0w"))
	assert.Equal(t, true, Valid("https://youtube.com/watch?v=Xq7z6WpeB0w"))
	assert.Equal(t, true, Valid("http://www.youtube.com/watch?v=Xq7z6WpeB0w"))
	assert.Equal(t, true, Valid("http://youtube.com/watch?v=Xq7z6WpeB0w"))
	assert.Equal(t, true, Valid("www.youtube.com/watch?v=Xq7z6WpeB0w"))
	assert.Equal(t, true, Valid("youtube.com/watch?v=Xq7z6WpeB0w"))
	assert.Equal(t, true, Valid("http://youtu.be/Xq7z6WpeB0w"))
	assert.Equal(t, true, Valid("www.youtube.com/Xq7z6WpeB0w"))

	assert.Equal(t, false, Valid("https://www.example.com/watch?v=Xq7z6WpeB0w"))
	assert.Equal(t, false, Valid("http://youtube/Xq7z6WpeB0w"))
	assert.Equal(t, false, Valid("youtube/Xq7z6WpeB0w"))
}
