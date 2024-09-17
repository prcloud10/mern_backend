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

Use the script for the cuda image creation

```sh
cd scripts/k3d/kuda
./build.sh
k3d cluster create <clustername> --image=$IMAGE --gpus=1 ...........
```



## Deploy Multi-Node Cluster

Execute scripts create.sh, delete.sh, start.sh and stop.sh to manage creation of the cluster

There is a special cluster config file (cluster.cuda.loca.yml) for the cuda version.

Check scripts and config files for detailed info about cluster creation.


