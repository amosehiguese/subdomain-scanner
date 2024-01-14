variable "gcp_project_id" {
    type = string
    description = "The GCP project ID to apply this config to"
}

variable "name" {
    type = string
    description = "Name given to the new GKE cluster"
    default = "subd"
}

variable "region" {
    type = string
    description = "Region of the new GKE cluster"
    default = "us-central1"
}

variable "namespace" {
    type = string
    description = "Kubernetes namespace in which the subd resources are to be deployed"
    default = "subd-prod"
}

