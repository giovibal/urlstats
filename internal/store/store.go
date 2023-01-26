package store

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type DB struct {
	urlStatsMap map[string]*UrlStats
	mu          sync.RWMutex
}

func New() *DB {
	return &DB{
		urlStatsMap: make(map[string]*UrlStats),
	}
}

func (db *DB) ExistsUrl(url string) bool {
	db.mu.RLock()
	defer db.mu.RUnlock()

	_, found := db.urlStatsMap[url]
	return found
}

func (db *DB) AddUrl(url string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	u, found := db.urlStatsMap[url]
	if !found {
		db.urlStatsMap[url] = &UrlStats{
			Url:       url,
			HitCount:  1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		return
	}
	// increment the url submission counter
	u.HitCount = u.HitCount + 1
	u.UpdatedAt = time.Now()
}

func (db *DB) IncrementUrlCounter(url string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	u, found := db.urlStatsMap[url]
	if !found {
		return fmt.Errorf("error incrementing url counter, url not found: %s", url)
	}
	// increment the url submission counter
	u.HitCount = u.HitCount + 1
	u.UpdatedAt = time.Now()
	return nil
}

func (db *DB) GetUrlList(orderBy []OrderBy, limit int) []UrlStats {
	db.mu.RLock()
	defer db.mu.RUnlock()

	list := []UrlStats{}

	for _, item := range db.urlStatsMap {
		list = append(list, *item)
	}

	// todo: sort slice (using params)

	for _, o := range orderBy {
		switch o {
		case OrderByCreatedAt:
			sort.Slice(list[:], func(i, j int) bool {
				return list[i].CreatedAt.UnixNano() < list[j].CreatedAt.UnixNano()
			})

		case OrderByCreatedAtDesc:
			sort.Slice(list[:], func(i, j int) bool {
				return list[i].CreatedAt.UnixNano() > list[j].CreatedAt.UnixNano()
			})

		case OrderByHitCount:
			sort.Slice(list[:], func(i, j int) bool {
				return list[i].HitCount < list[j].HitCount
			})

		case OrderByHitCountDesc:
			sort.Slice(list[:], func(i, j int) bool {
				return list[i].HitCount > list[j].HitCount
			})

		case OrderByAvgBytes:
			sort.Slice(list[:], func(i, j int) bool {
				return list[i].AvgBytes < list[j].AvgBytes
			})

		case OrderByAvgBytesDesc:
			sort.Slice(list[:], func(i, j int) bool {
				return list[i].AvgBytes > list[j].AvgBytes
			})

		default:
			sort.Slice(list[:], func(i, j int) bool {
				return list[i].CreatedAt.UnixNano() < list[j].CreatedAt.UnixNano()
			})
		}
	}

	// todo: limit to first N elements (using limit param)

	if limit < len(list) {
		return list[:limit]
	}
	return list
}

func (db *DB) Get10MostSubmittedUrls() []UrlStats {
	return db.GetUrlList([]OrderBy{OrderByHitCountDesc}, 10)
}

func (db *DB) UpdateUrlStats(url string, downloadSuccess bool, avgBytes int, downloadTime time.Duration) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	u, found := db.urlStatsMap[url]
	if !found {
		return fmt.Errorf("error updating url, url not found: %s", url)
	}

	u.UpdatedAt = time.Now()

	u.AvgBytes = uint64(avgBytes)
	if downloadSuccess {
		u.DownloadSuccessCount++
	} else {
		u.DownloadFailureCount++
	}
	u.DownloadTime = downloadTime

	return nil
}

func (db *DB) DeleteUrl(url string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, found := db.urlStatsMap[url]
	if !found {
		return fmt.Errorf("error deleting url, url not found: %s", url)
	}
	delete(db.urlStatsMap, url)
	return nil
}
