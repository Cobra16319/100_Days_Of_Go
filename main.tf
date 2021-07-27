terraform {
  backend "remote" {
    organization = "Cobra-Net"

    workspaces {
      name = "Go"
    }
  }
}
