run:
	go run cmd/main.go

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=donotshare -e POSTGRES_DB=predictionPlatformDb -d postgres:latest

docker-up:
	docker-compose up