#     
#     For testing purpose only. Production costs too $ 
#     The configuration for the `remote` backend.
      terraform {
        backend "remote" {
#         Name of your Terraform Cloud Org
          organization = "Cobra-Net"
#
#         # The name of the Terraform Cloud workspace to store Terraform state files in.
          workspaces {
            name = "100_Days_of_GO"
          }
        }
      }
#
#     # An example resource that does nothing.
      resource "null_resource" "example" {
        triggers = {
          value = "A example resource that does nothing!"
        }
      }
