terraform {
  required_providers {
    prefect2 = {
      version = "0.2"
      source  = "hashicorp.com/edu/prefect2"
    }
  }
}

provider "prefect2" {}