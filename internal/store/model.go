package store

import "time"

type UrlStats struct {
	Url                  string        `json:"url"`
	HitCount             uint64        `json:"hitCount"`
	DownloadSuccessCount uint64        `json:"downloadSuccessCount"`
	DownloadFailureCount uint64        `json:"downloadFailureCount"`
	AvgBytes             uint64        `json:"avgBytes"`
	DownloadTime         time.Duration `json:"downloadTime"`
	CreatedAt            time.Time     `json:"createdAt"`
	UpdatedAt            time.Time     `json:"updatedAt"`
}
