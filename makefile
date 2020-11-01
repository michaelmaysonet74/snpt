dev:
	go run cmd/snpt/main.go

docker-build:
	docker build -t snpt .

docker-run:
	docker run -d -p 9090:9090 snpt