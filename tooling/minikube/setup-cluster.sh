#!/bin/bash

GLOO_VERSION=$1
KNATIVE_VERSION=$2
EVENTING_VERSION=$3

# Setup Docker credentials to access private docker repos
echo "\n### Setup Docker credentials to access private repos ..."
if ! [[ -f tooling/docker/config.json ]]; then
  echo "Docker config file doesn't exist..."
  exit 1
fi

kubectl create secret generic dockercreds \
    --from-file=.dockerconfigjson=tooling/docker/config.json \
    --type=kubernetes.io/dockerconfigjson

kubectl patch serviceaccount default -p '{"imagePullSecrets": [{"name": "dockercreds"}]}'

# Install Knative Serving and Eventing via Gloo Cli
echo "\n### Installing Knative Serving & Eventing ..."
./gloo/binaries/glooctl-$GLOO_VERSION install knative -k --install-knative-version="$KNATIVE_VERSION" -e --install-eventing-version="$EVENTING_VERSION" -v

# Install Gloo Gateway Proxy
echo "\n### Installing Gloo Gateway Proxy..."
./gloo/binaries/glooctl-$GLOO_VERSION install gateway -v

kubectl label namespace default discovery.solo.io/function_discovery=enabled
kubectl label namespace transforms-demo discovery.solo.io/function_discovery=enabled

# Install and configure cert-manager
# echo "\n### Installing cert-manager..."
# kubectl apply -f infrastructure/k8s/cert-mgmr/00-crds.yaml --force=true

# kubectl apply -f infrastructure/k8s/cert-mgmr/cert-manager-v$3.yaml  --force=true
# sleep 30

# kubectl create secret generic $4 -n cert-manager --from-literal=access_key_id=$5 --from-literal=secret_access_key=$6

# kubectl apply -f infrastructure/k8s/cert-mgmr/production-issuer.yaml
# sleep 10

# kubectl apply -f infrastructure/k8s/cert-mgmr/certificate.yaml
