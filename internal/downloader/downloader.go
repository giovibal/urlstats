package downloader

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type DownloadHandler func(url string, isNew bool, downloadBytes int, elapsedTime time.Duration, err error)

type downloadJob struct {
	url   string
	isNew bool
}

type Downloader struct {
	urlsToDownload chan downloadJob
	maxWorkers     int
	queueSize      int
	handler        DownloadHandler
}

func NewDownloader(maxWorkers, queueSize int, handler DownloadHandler) *Downloader {
	return &Downloader{
		urlsToDownload: make(chan downloadJob, queueSize),
		maxWorkers:     maxWorkers,
		queueSize:      queueSize,
		handler:        handler,
	}
}

func (d *Downloader) Start() {
	for i := 0; i < d.maxWorkers; i++ {
		go d.StartWorker()
	}
}
func (d *Downloader) StartWorker() {
	for job := range d.urlsToDownload {
		url := job.url
		size, elapsedTime, err := DownloadUrl(url)
		d.handler(url, job.isNew, size, elapsedTime, err)
	}
}

func (d *Downloader) CheckNewUrl(url string) {
	job := downloadJob{url: url, isNew: true}
	d.urlsToDownload <- job
}
func (d *Downloader) CheckUrl(url string) {
	job := downloadJob{url: url, isNew: false}
	d.urlsToDownload <- job
}

// DownloadUrl download a url and return bytes downloaded and time to download
func DownloadUrl(url string) (int, time.Duration, error) {

	t1 := time.Now()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return -1, -1, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return -1, -1, fmt.Errorf("error http status: %d", resp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, -1, err
	}

	t2 := time.Now()

	elapsedTime := t2.Sub(t1)

	return len(bytes), elapsedTime, nil
}
