# K8i service to read k8 internals data


## Prerequisites

Install go:

```sh
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version
```


## Build and push to Registry

```sh
docker build . -t k8i
docker tag k8i:latest k3d-registry.localhost:12345/k8i:latest
docker push k3d-registry.localhost:12345/k8i:latest
```


## Deploy to Cluster



