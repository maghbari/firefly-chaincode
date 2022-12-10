# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0

ARG GO_VER=1.16
ARG ALPINE_VER=3.16.2

#FROM golang:${GO_VER}-alpine${ALPINE_VER}
FROM golang:1.16-alpine

# Set the Current Working Directory inside the container
WORKDIR /app
# Copy local code to the container image.
COPY  firefly.go go.mod go.sum ./
COPY chaincode/ ./chaincode/

RUN ls -la ./ \
    && GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build -o firefly firefly.go
    
RUN cp firefly /usr/local/bin/

EXPOSE 9999
CMD ["firefly"]
