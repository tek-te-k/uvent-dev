FROM golang:1.17.7-alpine

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update && apk upgrade && apk add git
RUN go get github.com/cosmtrek/air@v1.29.0

WORKDIR /app
COPY ./ ./

RUN go mod download

CMD ["air", "-c", ".air.toml"]
