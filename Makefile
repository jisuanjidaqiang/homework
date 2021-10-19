export tag=v1.0
root:
	export ROOT=github.com/jisuanjidaqiang/homework

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpserver container"
	docker build -t tbxxszq/homework:${tag} .

push: release
	echo "pushing cncamp/httpserver"
	docker push tbxxszq/homework:v1.0
