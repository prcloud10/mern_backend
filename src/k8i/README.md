# K8i service to read k8 internals data


## Prerequisites

Install go:

```sh
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version
```


## Build

```sh
docker build . -t app

```
Output:
Sending build context to Docker daemon   29.7kB
Step 1/13 : ARG GO_VERSION=1.23.0
Step 2/13 : FROM golang:${GO_VERSION}-alpine AS build
 ---> d0c638dc5c33
Step 3/13 : RUN apk add --no-cache git
 ---> Running in f4cfb62e3308




## Deploy to Cluster



