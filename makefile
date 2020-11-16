docker-build:
	docker build -t snpt .

dev:
	make docker-build && docker-compose up
