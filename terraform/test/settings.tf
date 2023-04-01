terraform {
  backend "local" {
    path = "./terraform.tfstate"
  }
  required_version = ">=1.4"
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~>3"
    }
  }
}

provider "azurerm" {
  features {}
}