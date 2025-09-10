terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "6.48.0"
    }
  }
}

provider "google" {
  project = var.project_id
  region  = var.region 
  zone    = var.zone
}

terraform {
  backend "gcs" {
    bucket  = "tf-bucket-074876"
    prefix  = "terraform/state"
  }
}

module "instances" {
  source = "./modules/instances"
}

module "storage" {
  source = "./modules/storage"
}

module "network" {
  source  = "terraform-google-modules/network/google"
  version = "10.0.0"
  # insert the 3 required variables here
  network_name = "tf-vpc-754385"
  routing_mode = "GLOBAL"
  project_id = var.project_id
  subnets = [
    {
      subnet_name = "subnet-01"
      subnet_region = var.region
      subnet_ip = "10.10.10.0/24"
    },
    {
      subnet_name = "subnet-02"
      subnet_region = var.region
      subnet_ip = "10.10.20.0/24"
    }
  ]
}

resource "google_compute_firewall" "tf-firewall" {
  name    = "tf-firewall"
  network = "tf-vpc-754385"

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }

  direction = "INGRESS"
  source_ranges = ["0.0.0.0/0"]
}
