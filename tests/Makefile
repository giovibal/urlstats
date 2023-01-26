test-add-urls:
	curl -X POST -d '{"url":"http://www.example.com"}' http://localhost:3000/

	curl -X POST -d '{"url":"http://www.example1.com"}' http://localhost:3000/

	curl -X POST -d '{"url":"http://www.example2.com"}' http://localhost:3000/

	curl -X POST -d '{"url":"http://www.example3.com"}' http://localhost:3000/

	curl -X POST -d '{"url":"http://example-ko.com"}' http://localhost:3000/

	curl -X POST -d '{"url":"http://www.example.com"}' http://localhost:3000/

	curl -X POST -d '{"url":"http://www.example.com"}' http://localhost:3000/

	curl -X POST -d '{"url":"http://www.example.com"}' http://localhost:3000/

	curl -X POST -d '{"url":"http://www.example1.com"}' http://localhost:3000/
	curl -X POST -d '{"url":"http://www.example2.com"}' http://localhost:3000/
	curl -X POST -d '{"url":"http://www.example3.com"}' http://localhost:3000/
	curl -X POST -d '{"url":"http://www.example4.com"}' http://localhost:3000/
	curl -X POST -d '{"url":"http://www.example5.com"}' http://localhost:3000/
	curl -X POST -d '{"url":"http://www.example6.com"}' http://localhost:3000/
	curl -X POST -d '{"url":"http://www.example7.com"}' http://localhost:3000/
	curl -X POST -d '{"url":"http://www.example8.com"}' http://localhost:3000/
	curl -X POST -d '{"url":"http://www.example9.com"}' http://localhost:3000/
	curl -X POST -d '{"url":"http://www.example10.com"}' http://localhost:3000/
	curl -X POST -d '{"url":"http://www.example11.com"}' http://localhost:3000/
	curl -X POST -d '{"url":"http://www.example12.com"}' http://localhost:3000/
	curl -X POST -d '{"url":"http://www.example13.com"}' http://localhost:3000/
	curl -X POST -d '{"url":"http://www.example14.com"}' http://localhost:3000/
	curl -X POST -d '{"url":"http://www.example15.com"}' http://localhost:3000/

	curl -X POST -d '{"url":"http://www.example10.com"}' http://localhost:3000/
	curl -X POST -d '{"url":"http://www.example11.com"}' http://localhost:3000/
	curl -X POST -d '{"url":"http://www.example12.com"}' http://localhost:3000/
	curl -X POST -d '{"url":"http://www.example13.com"}' http://localhost:3000/
	curl -X POST -d '{"url":"http://www.example14.com"}' http://localhost:3000/
	curl -X POST -d '{"url":"http://www.example15.com"}' http://localhost:3000/
	
	curl -X POST -d '{"url":"http://www.example13.com"}' http://localhost:3000/
	curl -X POST -d '{"url":"http://www.example14.com"}' http://localhost:3000/
	curl -X POST -d '{"url":"http://www.example15.com"}' http://localhost:3000/

	curl -X POST -d '{"url":"http://www.example.com"}' http://localhost:3000/

test-get-urls:
	# sorted from newest to oldest"
	curl -s http://localhost:3000/?orderBy=createdAtDesc | jq -c '.[]'

	# sorted from the smallest to the biggest
	curl -s http://localhost:3000/?orderBy=avgBytes | jq -c '.[]'
