FROM golang:alipine as build

ENV ENV=dev

WORKDIR /com

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

FROM alipine:latest as com

WORKDIR /com
COPY --chown=builder /com/yongzin /usr/bin/api
RUN chomd +x /usr/bin/api \
    && echo "Asia/Shanghai" > /etc/timezone
EXPOSE 3000

ENTRYPOINT [ "api" ]


