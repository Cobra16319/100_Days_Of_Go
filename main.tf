# Updated with guide to test CI
terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
    random = {
      source = "hashicorp/random"
    }
  }

  backend "remote" {
   organization = "Cobra-Net"

    workspaces {
      name = "gh-actions-demo"
    }
  }
}
