terraform {
  backend "remote" {
    organization = "Cobra-Net"

    workspaces {
      name = "100_Days_Of_Go"
    }
  }
}
