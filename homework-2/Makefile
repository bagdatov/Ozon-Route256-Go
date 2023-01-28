.PHONY: run

run:
	go run --race cmd/game-master/main.go && go run --race cmd/telegram-bot/main.go

mock:
	cd game-master/usecase && minimock -i Game,Database,Parser

test:
	go test -v ./game-master/repository/resty
	go test -v ./game-master/usecase
	go test -v ./game-master/delivery/grpc

up:
	docker-compose up

down:
	docker-compose down

proto:
	protoc -I api/. --go_out=api/. --go-grpc_out=api/. --grpc-gateway_out=api/. --openapiv2_out=api/swagger/. ./api/api.proto


.DEFAULT_GOAL := run