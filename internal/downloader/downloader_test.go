package downloader

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDownloadUrlSuccess(t *testing.T) {
	url := "http://www.example.com"
	size, elapsedTime, err := DownloadUrl(url)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, size, 0)
	assert.GreaterOrEqual(t, elapsedTime, time.Duration(0))
	log.Printf("%s: %d bytes in %s\n", url, size, elapsedTime)
}
func TestDownloadUrlSuccessHttps(t *testing.T) {
	url := "https://www.example.com"
	size, elapsedTime, err := DownloadUrl(url)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, size, 0)
	assert.GreaterOrEqual(t, elapsedTime, time.Duration(0))
	log.Printf("%s: %d bytes in %s\n", url, size, elapsedTime)
}

func TestDownloadUrlFailure(t *testing.T) {
	url := "http://www.example-ko.com"
	size, elapsedTime, err := DownloadUrl(url)
	assert.Error(t, err)
	assert.Equal(t, -1, size)
	assert.Equal(t, time.Duration(-1), elapsedTime)
}

func TestDownloadUrlFailureHttps(t *testing.T) {
	url := "https://www.example-ko.com"
	size, elapsedTime, err := DownloadUrl(url)
	assert.Error(t, err)
	assert.Equal(t, -1, size)
	assert.Equal(t, time.Duration(-1), elapsedTime)
}
