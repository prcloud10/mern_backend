#!/usr/bin/env bash
set -e

docker stop registry.localhost || true

k3d cluster stop local || true