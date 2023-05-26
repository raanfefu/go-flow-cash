build-golang:
	env GOOS=linux GOARCH=arm64 go build main.go

build-image:
	docker build -t lambda-main .

build:
	make build-golang
	make build-image

run:
	make build-golang && make build-image
	docker run --rm -p 9000:8080 lambda-main