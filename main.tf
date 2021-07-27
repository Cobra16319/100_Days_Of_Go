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
+   organization = "Cobra-Net"

    workspaces {
      name = "100_Days_Of_Go"
    }
  }
}
