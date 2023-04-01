module "thing1" {
  source = "./org/thing1"
  thing  = "thing1"
  guid   = var.guid
}

module "thing2" {
  source = "./org/thing2"
  thing  = "thing2"
  guid   = var.guid
}

variable "guid" {
  description = "GUID to append"
  type        = string
}