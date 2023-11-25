FROM golang:1.17

ENV GOPATH /usr/src/app/go
ARG dir=$GOPATH/src/github.com/mrauer

WORKDIR ${dir}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /usr/bin/gitdump

RUN rm -rf $GOPATH
