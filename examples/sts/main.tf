terraform {
  required_providers {
    minio = {
      source  = "aminueza/minio"
      version = ">= 2.0.2"
    }
  }
}

provider "minio" {
  minio_server   = var.minio_server
}

