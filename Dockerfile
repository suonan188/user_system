FROM golang:alpine as builder
LABEL MAINTAINER=suonan

ENV GO111MODULE=auto
ENV GOPROXY=https://goproxy.cn,direct
ENV ENV=dev

WORKDIR /app

COPY go.mod .
COPY go.sum .

#RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o api

FROM alpine:latest as app
ENV ENV=dev
WORKDIR /app

COPY --from=builder /app/api /usr/bin/api
ADD ./config /app/config

RUN chmod +x /usr/bin/api 
    #&& echo "Asia/Shanghai" > /etc/timezone
EXPOSE 3000

ENTRYPOINT [ "api" ]
# CMD ["/bin/bash", "/docker-build.sh"]
#CMD ["/bin/bash", "/app/build.sh"]