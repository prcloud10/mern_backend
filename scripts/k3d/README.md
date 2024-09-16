# K3d and Istio

## Prerequisites

Ensure `docker`, `k3d` , `helm` and `istioctl` installed.

```sh
wget -q -O - https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
curl -sL https://istio.io/downloadIstioctl | sh -
export PATH=$HOME/.istioctl/bin:$PATH
sudo curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
```

## Prerequisites for CUDA

Install NVIDIA Container Toolkit.

Build a new k3d image with the following changes: 

  1. Change the base images to nvidia/cuda:12.4.1-base-ubuntu22.04 so the NVIDIA Container Toolkit can be installed. 
  2. Add a manifest for the NVIDIA driver plugin for Kubernetes with an added RuntimeClass definition. 


Then build k3d cuda image and create the cluster with it.

```sh
cd scripts/k3d
./build.sh
k3d cluster create <clustername> --image=$IMAGE --gpus=1 ...........
```



## Deploy Multi-Node Cluster

First create ked-managed registry:
```sh
k3d registry create registry.localhost --port 5000
```

then create cluster with cuda image:

```sh
k3d cluster create cluster1 \
  --image=$IMAGE \
  --gpus=1 \
  --registry-use k3d-registry.localhost:5000 \
  --agents 1 \
  --port '32036:32036@server:0' \
  --port '8081:80@loadbalancer' \
  --api-port 6443 
  
```

## Probe New Cluster

```sh
docker push k3d-registry.localhost:12345/nginx:latest
docker ps --format 'table {{.ID}}\t{{.Image}}\t{{.Names}}\t{{.Ports}}'
kubectl get nodes
kubectl get ns
kubectl get pods -A
kubectl get services -A
```


## Install Istio

Istio install profiles.

Inspect profiles:

```sh
istioctl profile list # Will list available profiles
istioctl profile dump default # Will dump the default profile config
istioctl profile dump demo
```

Install demo profile:

```sh
istioctl install --set profile=demo -y
```

To enable the automatic injection of Envoy sidecar proxies, run the following:
(Otherwise you will need to do this manually when you deploy your applications)

```sh
kubectl label namespace default istio-injection=enabled
```

## Install Istio tools

Download and extract :

```sh
ISTIO_VERSION=1.10.0
ISTIO_URL=https://github.com/istio/istio/releases/download/$ISTIO_VERSION/istio-$ISTIO_VERSION-linux-amd64.tar.gz
curl -L $ISTIO_URL | tar xz
cd istio-$ISTIO_VERSION
```


## Install Tracing Utilities

```sh
kubectl apply -f samples/addons
kubectl rollout status deployment/kiali -n istio-system
```

If there are errors trying to install the addons, try running the command again.
There may be some timing issues which will be resolved when the command is run
again.

## Access the Kiali dashboard

```sh
istioctl dashboard kiali

for i in $(seq 1 100); do curl -so /dev/null http://localhost/productpage; done
```

## Uninstall and Cleanup

Uninstall Istio:

```sh
kubectl delete -f samples/addons
istioctl x uninstall --purge
kubectl delete namespace istio-system
kubectl label namespace default istio-injection-
```

Delete k3d/k3s cluster:

```sh
k3d cluster delete cluster1
```
