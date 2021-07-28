##This is just for CI/CD and for Terraform Workspace

terraform {
  backend "remote" {
    organization = "Cobra-Net"

    workspaces {
      name = "Go"
    }
  }
}
