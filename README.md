# URL Stats

The solution must be a single executable, a linux daemon, with an external yaml configuration file, containing relevant parameters.

## Features
- It should expose a minimal HTTP server. 
- The web server should serve the required REST APIs.
- The API should expose a method able to receive a request to add a website URL to an internal list of objects (example: http://www.example.com/).
- The API should expose a method able to retrieve the latest 50 URLs sent throuh the previous method (sorted from newest to oldest or from the smallest to the biggest, upon user request) and a counter that would show how many times that specific URL have been submitted to the API since the program started.
- Upon submission, each url should be downloaded (GET request) but not more than 3 downloads at the same time should be executed.  If the initial download of the page fails, throw the URL away otherwise it should be stored and used later.
- Create a background function executed every 60 seconds. This function must get the 10 most submitted/requested URLs from the ones that have been submitted and try to fetch the URL again, measuring the time it took to download it. All the download operations should happen in parallel with a concurrency factor of three - so no more than three GET requests should happen at the same time.
- Collect all the downloads times, successfull downloads counter and failed downloads counter and log them all on the stdout when the previous batch process completes.
- Ideally, the API should be ready to be used by a single page JS application running in another product.


# TODO

- use context to manage gracefull shutdown
- fix calculation of average bytes: add totBytes, and do `avgBytes = totBytes / downloadSuccessCount`
- aggregate store and downloader in an high level service (facade), and refactor main. Better readability and testing

