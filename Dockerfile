FROM --platform=linux/amd64 golang:1.23-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
ENV GOARCH amd64
ENV GOOS linux

# 替换为阿里云的镜像源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk add tzdata && apk add git

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o server ./cmd/app/main.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /build/server /app/server

CMD ["./server"]
