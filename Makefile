build:
	docker build -t gitdump:latest .

dev:
	docker build -t gitdump:latest . && docker run -it --rm -v ${CURDIR}:/usr/src/app/go/src/github.com/mrauer/gitdump gitdump:latest && docker exec -it gitdump:latest

release:
	env GOOS=darwin GOARCH=amd64 go build -i -o bin/osx/gitdump
	env GOOS=linux GOARCH=amd64 go build -i -o bin/linux/gitdump
	env GOOS=windows GOARCH=amd64 go build -i -o bin/windows/gitdump.exe
