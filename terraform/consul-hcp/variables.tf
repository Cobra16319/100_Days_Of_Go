variable "hvn_id" {
  description = "The ID of the HCP HVN."
  type        = string
  default     = "go-hvn"
}
variable "cluster_id" {
  description = "The ID of the HCP Consul cluster."
  type        = string
  default     = "go-hcp-consul"
}
variable "region" {
  description = "The region of the HCP HVN and Consul cluster."
  type        = string
  default     = "us-west-2"
}
variable "cloud_provider" {
  description = "The cloud provider of the HCP HVN and Consul cluster."
  type        = string
  default     = "aws"
}
variable "peering_id" {
  description = "The ID of the HCP peering connection."
  type        = string
  default     = "go-peering"
}
variable "route_id" {
  description = "The ID of the HCP HVN route."
  type        = string
  default     = "go-hvn-route"
}
