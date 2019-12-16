#!/bin/bash

## Check directory exists
DIR=$HOME/.minikube
if [[ ! -d "$DIR" ]]; then
  mkdir -p $HOME/.minikube/config
fi

## Create Minikube configuration
FILE=$HOME/.minikube/config/config.json
VERSION="$1"
if ! [[ -f "$FILE" ]]; then
  cat <<EOF > $FILE
  {
    "cpus": 4,
    "dashboard": true,
    "disk-size": "20g",
    "memory": 8192,
    "vm-driver": "hyperkit",
    "kubernetes-version": "$VERSION"
  }
EOF
else
  echo "Config exists... Skipping..."
fi
