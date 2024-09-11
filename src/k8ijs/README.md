# K8ijs service to read k8 internals data

JS Microservice for k8 internal api.


## Prerequisites

Install nvm, nodejs, npm:

```sh
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.0/install.sh | bash
nvm install 20
```

## Build and push to Registry

```sh
docker build . -t k8ijs
docker tag k8ijs:latest k3d-registry.localhost:12345/k8ijs:v0.6
docker push k3d-registry.localhost:12345/k8ijs:v0.6
```

## Give permission RBAC inside cluster or check if exists

```sh
kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=default:default
```


## Deploy to Cluster and test api

```sh
kubectl apply -f deploy.yml
curl http://localhost:8081/k8ijs/api
```

