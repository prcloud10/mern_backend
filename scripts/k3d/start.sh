#!/usr/bin/env bash
set -e

docker start registry.localhost || true

k3d cluster start local || true