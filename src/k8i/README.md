# K8i service to read k8 internals data

Golang Microservice for k8 internal api.


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
docker tag k8i:latest k3d-registry.localhost:12345/k8i:v0.1
docker push k3d-registry.localhost:12345/k8i:v0.1
```

## Give permission RBAC inside cluster

```sh
kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=default:default
```


## Deploy to Cluster and test api

```sh
kubectl apply -f deploy.yml
curl http://localhost:8081/k8i/api
```

