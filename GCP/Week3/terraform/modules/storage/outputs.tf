resource "google_storage_bucket" "tf-bucket-074876" {
  name          = "tf-bucket-074876"
  location      = "US"
  force_destroy = true

  uniform_bucket_level_access = true
}