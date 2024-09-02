# K3d and Istio

## Prerequisites

Ensure `docker`, `k3d` , `helm` and `istioctl` installed.

```sh
wget -q -O - https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
curl -sL https://istio.io/downloadIstioctl | sh -
export PATH=$HOME/.istioctl/bin:$PATH
sudo curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
```

## Deploy Multi-Node Cluster

```sh
k3d cluster create bob --servers 1 --agents 3 \
  --port 9443:443@loadbalancer \
  --port 80:80@loadbalancer \
  --port '32036:32036@server[0]' \
  --api-port 6443 --k3s-server-arg '--no-deploy=traefik'
```

## Probe New Cluster

```sh
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

## Deploy Demo Distributed System

Download and extract samples:

```sh
ISTIO_VERSION=1.10.0
ISTIO_URL=https://github.com/istio/istio/releases/download/$ISTIO_VERSION/istio-$ISTIO_VERSION-linux-amd64.tar.gz
curl -L $ISTIO_URL | tar xz
cd istio-$ISTIO_VERSION
```

```sh
kubectl apply -f samples/bookinfo/platform/kube/bookinfo.yaml
kubectl get services
kubectl get pods
```

## Verify Demo System

```sh
kubectl exec "$(kubectl get pod -l app=ratings \
  -o jsonpath='{.items[0].metadata.name}')" -c ratings \
  -- curl -sS productpage:9080/productpage | grep -o "<title>.*</title>"
```

## Open Outside Traffic

```sh
kubectl apply -f samples/bookinfo/networking/bookinfo-gateway.yaml
istioctl analyze
```

Open in browser:

```sh
open http://localhost/productpage
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
kubectl delete -f samples/bookinfo/networking/bookinfo-gateway.yaml
kubectl delete -f samples/bookinfo/platform/kube/bookinfo.yaml
istioctl x uninstall --purge
kubectl delete namespace istio-system
kubectl label namespace default istio-injection-
```

Delete k3d/k3s cluster:

```sh
k3d cluster delete bob
```
