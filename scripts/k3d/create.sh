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

echo "Creating local cluster..."
k3d cluster create local \
  --config ./cluster.local.yml 
#  1>"${log_dir}/k3s.log" \
#  2>"${log_dir}/k3s.err.log"
echo " Done!"

wait_for 2m

echo "Installing Istio"
helm repo add istio https://istio-release.storage.googleapis.com/charts
helm repo update
helm install istio-base istio/base -n istio-system --create-namespace --set defaultRevision=default
helm install istiod istio/istiod -n istio-system --wait
echo " Done!"


echo "Installing K8 Dashboard"
helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard/
helm upgrade --install kubernetes-dashboard kubernetes-dashboard/kubernetes-dashboard --create-namespace --namespace kubernetes-dashboard
echo " Done!"


echo "Installing ArgoCD"
helm repo add argo https://argoproj.github.io/argo-helm
helm install argocd argo/argo-cd --create-namespace --namespace argocd
echo " Done!"


echo
echo "Cluster is started up!"
echo
echo "The following Services are provided out of the box:"
echo "    Traefik loadbalancer"
echo "    Istio"
echo "    ArgoCD ( user:admin and password:<get from argocd-initial-admin-secret> )"
echo 
echo "The system provides a registry at"
echo "    localhost:5000"
echo
echo "If you want to deploy custom images, push them into this registry."
echo "To deploy images from this registry, reference them as registry.localhost:5000/<image-name>:<image-tag>"
echo
echo "The kubernetes configuration has been written"
