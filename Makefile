default: build

build:
	go build -o ./build/urlstats ./cmd/urlstats

clean:
	rm -Rf ./build

run:
	go run ./cmd/urlstats

test:
	make -C tests test-add-urls
	sleep 10
	make -C tests test-get-urls
