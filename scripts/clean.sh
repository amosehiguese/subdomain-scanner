#!/bin/bash

set -euo pipefail

REGION=$1
CLUSTER_NAME=$2
LOCAL=${3:-true}

delete_resources() {
  echo "Deleting Kibana..."
  kubectl delete kibana quickstart || true

  echo "Deleting Fluent Bit..."
  helm uninstall fluent-bit || true

  echo "Deleting Elasticsearch..."
  kubectl delete elasticsearch quickstart || true

  echo "Deleting Elasticsearch CRDs and Operator..."
  kubectl delete -f -f https://download.elastic.co/downloads/eck/2.5.0/operator.yaml || true
  kubectl delete -f https://download.elastic.co/downloads/eck/2.5.0/crds.yaml || true

  echo "Deleting Subdomain Chart"
  helm uninstall subd || true

  echo "✅ K8S resources deleted."
}

delete_kind_cluster() {
  echo "Deleting kind cluster..."
  kind delete cluster --name kind || true
  echo "✅ Kind cluster deleted"
}

delete_aws_cluster() {
  echo "Deleting AWS EKS cluster..."
  aws eks delete-cluster --name "$CLUSTER_NAME" --region "$REGION" || true
  echo "✅ AWS EKS cluster deleted."
}

if $LOCAL; then
  delete_resources
  delete_kind_cluster
else
  delete_resources
  delete_aws_cluster
fi

echo "✅ Cleanup complete."
