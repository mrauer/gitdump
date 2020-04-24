dev: build run
osx: build mac

build:
	docker build -t gitdump:latest .

run:
	docker build -t gitdump:latest . && docker run -it --rm -v ${CURDIR}:/usr/src/app/go/src/github.com/mrauer/gitdump gitdump:latest && docker exec -it gitdump:latest

mac:
	docker run --rm -w /usr/src/app/go/src/github.com/mrauer/gitdump -v ${CURDIR}:/usr/src/app/go/src/github.com/mrauer/gitdump gitdump:latest env GOOS=darwin GOARCH=amd64 go build -i -o gitdump && chmod +x gitdump
