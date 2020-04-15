dev:
	docker build -t gitdump:latest . && docker run -it --rm -v ${CURDIR}:/usr/src/app gitdump:latest && docker exec -it gitdump:latest

osx:
	docker run --rm -v ${CURDIR}:/usr/src/app -w /usr/src/app golang:1.13 env GOOS=darwin GOARCH=amd64 go build -i -o gitdump && chmod +x gitdump
