package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/giovibal/urlstats/internal/downloader"
	"github.com/giovibal/urlstats/internal/store"
)

const (
	maxParallelDownloads int = 3  // todo: get from config
	limit                int = 50 // todo: get from config
)

func main() {

	// Init an 'in memory' DB
	db := store.New()

	// Start downloader job
	downloaderJob := downloader.NewDownloader(maxParallelDownloads, func(url string, isNew bool, downloadBytes int, elapsedTime time.Duration, err error) {
		downloadSuccess := err == nil
		if err != nil {
			log.Printf("error downloading url: %s, %s", url, err)
		}

		if isNew {
			if downloadSuccess {
				db.AddUrl(url)
				updErr := db.UpdateUrlStats(url, downloadSuccess, downloadBytes, elapsedTime)
				if updErr != nil {
					log.Println(err)
				}
			}
			// don't add url to DB if download fails
			return
		}

		// Collect all the downloads times, successfull downloads counter and failed downloads counter ...
		updErr := db.UpdateUrlStats(url, downloadSuccess, downloadBytes, elapsedTime)
		if updErr != nil {
			log.Println(updErr)
		}
		// ... and log them all on the stdout when the previous batch process completes.
		log.Printf("check url: %s, download time: %s, download bytes %d, download error %s", url, elapsedTime, downloadBytes, err)
	})
	downloaderJob.Start()

	// Create a background function executed every 60 seconds.
	// This function must get the 10 most submitted/requested URLs from the ones that have been submitted
	// and try to fetch the URL again, measuring the time it took to download it.
	// All the download operations should happen in parallel with a concurrency factor of three - so no more than three GET requests should happen at the same time.
	ticker := time.NewTicker(60 * time.Second)
	go func() {
		for {
			select {
			// case <-ctx.Done():
			// 	return
			case t := <-ticker.C:
				log.Println("Start check '10 most submitted/requested URLs' job", t)

				urlsToCheck := db.Get10MostSubmittedUrls()
				for _, item := range urlsToCheck {
					downloaderJob.CheckUrl(item.Url)
				}

			}
		}
	}()

	// Serve rest APIs
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Request timeout
	r.Use(middleware.Timeout(60 * time.Second))

	// CORS
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		orderByReqParam := r.URL.Query().Get("orderBy")
		orderByArray := strings.Split(orderByReqParam, ",")
		orderBy := store.OrderByFromStringArray(orderByArray)

		list := db.GetUrlList(orderBy, limit)

		b, err := json.Marshal(list)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(b)
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		var req addUrlRequest
		json.Unmarshal(b, &req)

		url := req.Url

		// downloadBytes, elapsedTime, err := downloader.DownloadUrl(url)
		// if err != nil {
		// 	http.Error(w, err.Error(), 500)
		// 	return
		// }
		// db.AddUrl(url)

		// downloadSuccess := err != nil
		// avgBytes := downloadBytes
		// updErr := db.UpdateUrlStats(url, downloadSuccess, avgBytes, elapsedTime)
		// if updErr != nil {
		// 	log.Println(err)
		// }
		if db.ExistsUrl(url) {
			db.IncrementUrlCounter(url)
		} else {
			downloaderJob.CheckNewUrl(url)
		}

		w.Write(b)
	})

	http.ListenAndServe(":3000", r)
}

type addUrlRequest struct {
	Url string `json:"url"`
}
