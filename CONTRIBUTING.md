# How to Contribute

Thank you so much for your interest in contributing to Subdomain Enumeration Tool.

## Development Guide

This doc explains how to build and run the Subdomain Scanner source code locally using the `skaffold` command-line tool.

## Prerequisites

- [Docker for Desktop](https://www.docker.com/products/docker-desktop) for windows
- [Docker Engine](https://docs.docker.com/engine/install/) for linux
- [Kind](https://kind.sigs.k8s.io/) (optional Local Cluster 2)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- Clone the repository.
  ```sh
  git clone https://github.com/amosehiguese/subdomain-scanner.git
  cd subdomain-scanner/
  ```

## Run on a Local Cluster

1. Launch the appliction
  ```
    chmod +x ./scripts/run.sh
    ./scripts/run.sh

  ```
2. Run `kubectl get pods` to verify the Pods are ready and running.

3. Run `kubectl port-forward pod/<Pod_name> 8080:8080` to forward a port to the frontend.

4. Navigate to `localhost:8080` to access the web frontend.

## Cleanup

To clean up the deployed resources. Run

```
chmod +x ./scripts/clean.sh
./scripts/clean.sh
```
