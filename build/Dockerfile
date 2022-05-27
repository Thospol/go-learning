FROM golang:1.18.1-alpine3.15 as builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN apk update && apk upgrade && \
  apk add --no-cache ca-certificates git openssh
  

RUN mkdir /api
WORKDIR /api
ADD . /api

RUN go mod download
RUN go build -o api

FROM alpine:3.15

RUN apk update && apk upgrade && \
  apk add --no-cache ca-certificates tzdata

COPY --from=builder /api/api .

ADD /configs /configs
ADD /assets /assets

ENV TZ=Asia/Bangkok

EXPOSE 8000

ENTRYPOINT ["/api"]