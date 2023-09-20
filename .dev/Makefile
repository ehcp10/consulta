teste:
	mkdir -p bin
	go test -short -coverprofile=bin/cov.out `go list ./... | grep -v vendor`
up:
	docker compose up -d

down:
	docker compose down

stop: 
	docker compose stop

run:
	go run main.go