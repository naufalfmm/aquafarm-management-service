FROM golang:1.19 as builder

# install xz
RUN apt-get update && apt-get install -y \
    xz-utils \
&& rm -rf /var/lib/apt/lists/*
# install UPX
ADD https://github.com/upx/upx/releases/download/v3.94/upx-3.94-amd64_linux.tar.xz /usr/local
RUN xz -d -c /usr/local/upx-3.94-amd64_linux.tar.xz | \
    tar -xOf - upx-3.94-amd64_linux/upx > /bin/upx && \
    chmod a+x /bin/upx

WORKDIR /go/src/github.com/naufalfmm/aquafarm-management-service
COPY go.mod go.sum ./

RUN GO111MODULE=on go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o naufalfmm/aquafarm-management-service

RUN strip --strip-unneeded naufalfmm/aquafarm-management-service
RUN upx naufalfmm/aquafarm-management-service

FROM alpine:latest
RUN apk update && apk add --no-cache tzdata
RUN apk --no-cache add ca-certificates

WORKDIR /usr/src

COPY --from=builder /go/src/github.com/naufalfmm/aquafarm-management-service .

CMD ["./naufalfmm/aquafarm-management-service"]
