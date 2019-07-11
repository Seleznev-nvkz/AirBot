ARG GO_VERSION=1.12.6

FROM golang:${GO_VERSION}-alpine3.9 AS builder

ARG GO111MODULE=on
ARG CGO_ENABLED=0
ARG GOOS=linux

RUN apk add --no-cache \
        ca-certificates \
        git
WORKDIR /go/src/AirBot/

COPY go.* /go/src/AirBot/
RUN go mod download

COPY *.go /go/src/AirBot/
RUN go build -a -ldflags="-s -w" -installsuffix cgo -o /bot .


FROM kudato/baseimage:alpine3.9
ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
WORKDIR /usr/bin/
COPY --from=builder /bot /usr/bin/bot
CMD [ "/usr/bin/bot" ]
