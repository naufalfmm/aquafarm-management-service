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

WORKDIR /go/src/github.com/naufalfmm/aquafarm-management-service-migration
COPY go.mod go.sum ./

# install modules
RUN GO111MODULE=on go mod download

COPY /migrations/. ./migrations/.
COPY /consts/. ./consts/.
COPY /utils/. ./utils/.
COPY .env .env