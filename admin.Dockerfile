# 构建
FROM golang:1.14-alpine as builder
WORKDIR /application
ENV GOPROXY=https://goproxy.cn
COPY ./go/ ./
RUN go mod download && \
sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
apk add tzdata
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o admin

# 打包
FROM alpine as runner
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /application/admin /application/
WORKDIR /application
CMD ["./admin"]
