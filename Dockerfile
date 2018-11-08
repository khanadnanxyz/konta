FROM golang:onbuild
EXPOSE 8080

RUN apk add --no-cache --update git
ENV GOPATH=/go

RUN go get -u github.com/golang/dep/cmd/dep
COPY . $GOPATH/src/github.com/khanadnanxyz/konta
WORKDIR $GOPATH/src/github.com/khanadnanxyz/konta

RUN ./build.sh