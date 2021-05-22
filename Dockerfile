FROM golang:1.14

ENV GOPATH /usr/src/app/go
ARG dir=$GOPATH/src/github.com/mrauer
WORKDIR ${dir}

COPY go.mod .
COPY go.sum .
RUN go mod download

WORKDIR $GOPATH/src/github.com/mrauer/gitdump
COPY . .
