.PHONY: build_greet docker_greet run_greet swag_greet

build_greet:
	go build -o greet cmd/greet/main.go

docker_greet:
	docker build -t greet:latest -f build/greet/Dockerfile .  --network=host

run_greet:
	go run cmd/greet/main.go -c cmd/greet/res

swag_greet:
	swag init -g internal/irisscaffold/greet/route.go --output ./docs/greet --parseInternal true --pd true