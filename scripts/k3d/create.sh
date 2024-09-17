#!/usr/bin/env bash
set -e

function wait_for() {
  timeout \
    --foreground \
    --kill-after "${1}" \
    "${1}" \
    bash -c \
      "while :
        printf '.'
      do
        ${2} 1>/dev/null 2>/dev/null && break
        sleep 5
      done"
}


log_dir="logs"
rm -rf "${log_dir:?}/"
mkdir -p "${log_dir}"

echo "creating local cluster..."
k3d cluster create local \
  --config ./cluster.local.yml 
#  1>"${log_dir}/k3s.log" \
#  2>"${log_dir}/k3s.err.log"
echo " Done!"

wait_for 2m

#echo "installing Istio"
#export PATH=$HOME/.istioctl/bin:$PATH
#istioctl profile list 
#istioctl profile dump default 
#istioctl profile dump demo
#istioctl install --set profile=demo -y
#kubectl label namespace default istio-injection=enabled
#ISTIO_VERSION=1.10.0
#ISTIO_URL=https://github.com/istio/istio/releases/download/$ISTIO_VERSION/istio-$ISTIO_VERSION-linux-amd64.tar.gz
#curl -L $ISTIO_URL | tar xz
#cd istio-$ISTIO_VERSION
#wait_for 2m
#kubectl apply -f samples/addons
#kubectl rollout status deployment/kiali -n istio-system
#rm -rf ../istio-$ISTIO_VERSION
#echo " Done!"

echo "installing ArgoCD"
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
echo " Done!"


echo
echo "Cluster is started up!"
echo
echo "The following Services are provided out of the box:"
echo "    traefik loadbalancer"
echo "    Istio"
echo "    Kiali"
echo "    Prometheus"
echo "    Grafana"
echo 
echo "The system provides a registry at"
echo "    localhost:5000"
echo
echo "If you want to deploy custom images, push them into this registry."
echo "To deploy images from this registry, reference them as registry.localhost:5000/<image-name>:<image-tag>"
echo
echo "The kubernetes configuration has been written"
