FROM golang:1.13

ENV GOPATH /usr/src/app/go
ARG dir=$GOPATH/src/github.com/mrauer
WORKDIR ${dir}
COPY ./Gopkg.toml ${dir}/
COPY ./Gopkg.lock ${dir}/
RUN mkdir -p /usr/src/app/go/bin && \ 
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
    $GOPATH/bin/dep ensure --vendor-only -v

WORKDIR $GOPATH/src/github.com/mrauer/gitdump
COPY . .
