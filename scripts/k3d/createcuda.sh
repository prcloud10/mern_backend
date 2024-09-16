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

printf "creating local cluster..."

k3d cluster create local \
  --config ./cluster.cuda.local.yml \
  1>"${log_dir}/k3s.log" \
  2>"${log_dir}/k3s.err.log"
echo " Done!"


echo
echo "Cluster is started up!"
echo
echo "The following Services are provided out of the box:"
echo "    https://grafana.k3d.localhost (username: 'admin', password: 'prom-operator')"
echo "    https://prometheus.k3d.localhost"
echo "    https://traefik.k3d.localhost"
echo "    https://local.registry.k3d.localhost"
echo
echo "The system provides a registry at"
echo "    localhost:5000"
echo
echo "If you want to deploy custom images, push them into this registry."
echo "To deploy images from this registry, reference them as registry.localhost:5000/<image-name>:<image-tag>"
echo
echo "The kubernetes configuration has been written"
