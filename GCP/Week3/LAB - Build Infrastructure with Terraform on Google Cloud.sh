ROOT_DIR=$(pwd)

touch main.tf variables.tf

mkdir -p modules/instances && cd modules/instances
touch instances.tf outputs.tf variables.tf

cd $ROOT_DIR
mkdir -p modules/storage && cd modules/storage
touch storage.tf outputs.tf variables.tf

cd $ROOT_DIR


# Import VMs to state
terraform import google_compute_instance.tf-instance-1 projects/qwiklabs-gcp-01-8a54b32c0625/zones/us-west1-a/instances/tf-instance-1
terraform import google_compute_instance.tf-instance-2 projects/qwiklabs-gcp-01-8a54b32c0625/zones/us-west1-a/instances/tf-instance-2