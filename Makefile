
.PHONY: run up clean 

run:
	go run server.go

up:	
	rm -f grm-api-linux
	GOOS=linux GOARCH=amd64 go build -o grm-api-linux
	docker-compose up --build

clean:
	docker-compose down
	rm -rf grm-api
	rm -rf grm-api-linux
