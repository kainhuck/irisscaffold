.PHONY: build_greet docker_greet

build_greet:
	go build -o greet cmd/greet/main.go

docker_greet:
	docker build -t greet:latest -f build/greet/Dockerfile .  --network=host