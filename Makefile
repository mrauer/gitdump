osx:
	docker run --rm -v ${CURDIR}:/usr/src/app -w /usr/src/app golang:1.13 env GOOS=darwin GOARCH=amd64 go build -i -o gitdump && chmod +x gitdump
