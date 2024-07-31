#!/bin/bash

set -euo pipefail

check_and_install() {
    local cmd=$1
    local install_cmd=$2
    if ! command -v $cmd &> /dev/null; then
        echo "$cmd not found. Installing..."
        eval $install_cmd
        if ! command -v $cmd &> /dev/null; then
            echo "Failed to install $cmd. Exiting."
            exit 1
        fi
        echo "✅ $cmd installed"
    else
        echo "✅ $cmd is already installed."
    fi
}

check_and_install kind "curl -Lo ./kind https://kind.sigs.k8s.io/dl/latest/kind-linux-amd64 && chmod +x ./kind && sudo mv ./kind /usr/local/bin/kind"
check_and_install docker "curl -fsSL https://get.docker.com -o get-docker.sh && sh get-docker.sh && rm get-docker.sh"
check_and_install helm "curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash"
check_and_install kubectl "sudo apt-get install -yqq kubectl git"

CONTAINER_REGISTRY="ghcr.io"
LOCAL=true

load_images() {
  SERVICES=("apiquerysvc" "dnsresolvesvc" "portscansvc" "frontend")
  for service in "${SERVICES[@]}"; do
      cd src/$service || { echo "Service path $service not found. Exiting."; exit 1; }

      # Extract service name and version from directory or any other method you prefer
      service_name=$(basename "$service")
      version=$(git describe --match 'v[0-9]*' --tags --always)

      # Build Docker image
      docker build -t "$service_name:$version" .
      if [ $? -ne 0 ]; then
          echo "Failed to build Docker image for $service_name. Exiting."
          exit 1
      fi

      if [ "$LOCAL" = true ]; then
          # Load Docker image into kind cluster
          kind load docker-image "$service_name:$version"
      else
          # Tag and push Docker image to the registry
          docker tag "$service_name:$version" "$CONTAINER_REGISTRY/$service_name:$version"
          docker push "$CONTAINER_REGISTRY/$service_name:$version"
      fi

      cd - || exit
  done
}

is_kind_cluster_running() {
  kind get clusters | grep -q 'kind'
  return $?
}

create_kind_cluster() {
  echo "Creating a kind cluster with one control node and three worker nodes..."
  cat <<EOF | kind create cluster --name kind --config=-
  kind: Cluster
  apiVersion: kind.x-k8s.io/v1alpha4
  nodes:
    - role: control-plane
    - role: worker
    - role: worker
    - role: worker
EOF
  echo "✅ Kind cluster creation initiated."
}

check_all_nodes_running() {
  echo "Checking if all nodes are up and running..."
  while true; do
    ready_nodes=$(kubectl get nodes --no-headers | grep ' Ready ' | wc -1)
    if [ "$ready_nodes" -gt 0]; then
      echo "✅ All nodes are up and running."
      break
    else
      echo "Waiting for all nodes to be up and running..."
      sleep 5
    fi
  done
}

create_metrics_resources() {
  echo "Creating Kubernetes resources for Jaeger All In One, OpenTelemetry Collector, and Prometheus..."

  helm repo add huseyinbabal https://huseyinbabal.github.io/charts
  helm install subd-jaeger huseyinbabal/jaeger -n jaeger --create-namespace

  echo "✅ Successful installation of Metrics Resources."
}

create_logging_resources() {
  echo "Creating Kubernetes resources for Elasticsearch, Fluent Bit, and Kibana..."

  # Create CRDs and Operator for Elasticsearch
  kubectl create -f https://download.elastic.co/downloads/eck/2.5.0/crds.yaml
  kubectl apply -f https://download.elastic.co/downloads/eck/2.5.0/operator.yaml

  echo "✅ Elasticsearch CRDs and operator created."

  while true; do
    operator_status=$(kubectl get pods -n elastic -l control-plane=elastic-operator --no-headers | grep ' Running ' | wc -1)
    if [ "$operator_status" -gt 0]; then
      echo "✅ Elasticsearch operator is ready."
      break
    else
      echo "Waiting for the Elasticsearch operator to be ready..."
      sleep 10
    fi
  done

    # Create Elasticsearch cluster
  cat <<EOF | kubectl apply -f -
  apiVersion: elasticsearch.k8s.elastic.co/v1
  kind: Elasticsearch
  metadata:
    name: quickstart
  spec:
    version: 7.16.2
    nodeSets:
    - name: default
      count: 3
      config:
        node.store.allow_mmap: false
EOF
  echo "✅ Elasticsearch cluster creation request submitted."

  # Wait for the Elasticsearch cluster to be ready
  echo "Waiting for the Elasticsearch cluster to be ready..."
  while true; do
    es_status=$(kubectl get elasticsearch quickstart -o jsonpath='{.status.health}' 2>/dev/null)
    if [ "$es_status" == "green" ]; then
      echo "✅ Elasticsearch cluster is ready"
      break
    else
      echo "Waiting for the Elasticsearch cluster to be ready..."
      sleep 10
    fi
  done

  PASSWORD=$(kubectl get secret quickstart-es-elastic-user -o=jsonpath='{.data.elastic}' | base64 --decode)
  echo "✅ Elasticsearch password retrieved."

  echo "Setting up Fluent Bit..."
  helm repo add fluent-bit https://fluent.github.io/helm-charts
  helm repo update

  # Create Fluent Bit Configuration with Elasticsearch password
  cat <<EOF | helm install fluent-bit fluent/fluent-bit -f -
  config:
    outputs: |
      [OUTPUT]
        Name es
        Match kube.*
        Host quickstart-es-http
        HTTP_User elastic
        HTTP_Password $PASSWORD
        tls On
        tls.verify Off
        Logstash_Format On
        Retry_Limit False
EOF

  echo "✅ Fluent Bit installed successfully"

  # Install Kibana
  echo "Installing Kibana..."
  cat <<EOF | kubectl apply -f -
  apiVersion: kibana.k8s.elastic.co/v1
  kind: Kibana
  metadata:
    name: quickstart
  spec:
    version: 7.16.2
    count: 1
    elasticsearchRef:
      name: quickstart
EOF
  echo "✅ Kibana installation initiated."
  echo "✅ Successfully installation of EFK stack."

}


REGION=$1
CLUSTER_NAME=$2

env_vars=("REGION" "CLUSTER_NAME")

check_env_vars() {
  for v in "${env_vars[@]}"; do
    if [ -z "${!v}" ]; then
      echo "Env Var $v is not set"
      exit 1
    else
      echo "Env var $v is set"
    fi
  done
}

set_up_aws() {
  check_env_vars
  # Check and install AWS CLI
  check_and_install aws "curl https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip -o awscliv2.zip && unzip awscliv2.zip && sudo ./aws/install"
  aws eks --region $REGION update-kubeconfig --name $CLUSTER_NAME
}


# Main
if "$LOCAL"; then
  echo "✅ Setting up local cluster"
  if is_kind_cluster_running; then
    echo "✅ Kind cluster up"
  else
    echo "Kind cluster not up"
    create_kind_cluster
  fi
else
  echo "✅ Setting up EKS"
  set_up_aws $1 $2
fi

load_images
check_all_nodes_running
create_metrics_resources
create_logging_resources

echo "Installing Subd Helm Chart..."
helm install subd ./deploy/subd

echo "Execution completed."
