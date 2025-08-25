REGION="us-central1"
ZONE="us-central1-c"
gcloud config set compute/region $REGION
gcloud config set compute/zone $ZONE

# Task1: Create dev VPC and subnets
DEV_VPC_NAME="griffin-dev-vpc"
gcloud compute networks create $DEV_VPC_NAME --subnet-mode=custom

DEV_VPC_SUBNET1_NAME="griffin-dev-wp"
DEV_VPC_SUBNET1_REGION="$REGION"
DEV_VPC_SUBNET1_RANGE="192.168.16.0/20"
gcloud compute networks subnets create $DEV_VPC_SUBNET1_NAME --network=$DEV_VPC_NAME --region=$DEV_VPC_SUBNET1_REGION --range=$DEV_VPC_SUBNET1_RANGE

DEV_VPC_SUBNET2_NAME="griffin-dev-mgmt"
DEV_VPC_SUBNET2_REGION="$REGION"
DEV_VPC_SUBNET2_RANGE="192.168.32.0/20"
gcloud compute networks subnets create $DEV_VPC_SUBNET2_NAME --network=$DEV_VPC_NAME --region=$DEV_VPC_SUBNET2_REGION --range=$DEV_VPC_SUBNET2_RANGE

DEV_VPC_SUBNET3_NAME=""
DEV_VPC_SUBNET3_REGION=""
DEV_VPC_SUBNET3_RANGE=""
gcloud compute networks subnets create $DEV_VPC_SUBNET3_NAME --network=$DEV_VPC_NAME --region=$DEV_VPC_SUBNET3_REGION --range=$DEV_VPC_SUBNET3_RANGE

# Task2: Create prod VPC and subnets
PROD_VPC_NAME="griffin-prod-vpc"
gcloud compute networks create $PROD_VPC_NAME --subnet-mode=custom

PROD_VPC_SUBNET1_NAME="griffin-prod-wp"
PROD_VPC_SUBNET1_REGION="$REGION"
PROD_VPC_SUBNET1_RANGE="192.168.48.0/20"
gcloud compute networks subnets create $PROD_VPC_SUBNET1_NAME --network=$PROD_VPC_NAME --region=$PROD_VPC_SUBNET1_REGION --range=$PROD_VPC_SUBNET1_RANGE

PROD_VPC_SUBNET2_NAME="griffin-prod-mgmt"
PROD_VPC_SUBNET2_REGION="$REGION"
PROD_VPC_SUBNET2_RANGE="192.168.64.0/20"
gcloud compute networks subnets create $PROD_VPC_SUBNET2_NAME --network=$PROD_VPC_NAME --region=$PROD_VPC_SUBNET2_REGION --range=$PROD_VPC_SUBNET2_RANGE

PROD_VPC_SUBNET3_NAME=""
PROD_VPC_SUBNET3_REGION=""
PROD_VPC_SUBNET3_RANGE=""
gcloud compute networks subnets create $PROD_VPC_SUBNET3_NAME --network=$PROD_VPC_NAME --region=$PROD_VPC_SUBNET3_REGION --range=$PROD_VPC_SUBNET3_RANGE


# Task3: Create bastion host
gcloud compute instances create bastion-host \
    --network-interface network=$DEV_VPC_NAME,subnet=$DEV_VPC_SUBNET2_NAME \
    --network-interface network=$PROD_VPC_NAME,subnet=$PROD_VPC_SUBNET2_NAME \
    --zone=$ZONE \
    --machine-type=e2-medium

gcloud compute firewall-rules create "${DEV_VPC_SUBNET2_NAME}-allow-icmp-ssh-rdp" --direction=INGRESS --priority=1000 --network=$DEV_VPC_NAME --action=ALLOW --rules=icmp,tcp:22,tcp:3389 --source-ranges=0.0.0.0/0
gcloud compute firewall-rules create "${PROD_VPC_SUBNET2_NAME}-allow-icmp-ssh-rdp" --direction=INGRESS --priority=1000 --network=$PROD_VPC_NAME --action=ALLOW --rules=icmp,tcp:22,tcp:3389 --source-ranges=0.0.0.0/0


# Task4: Create and configure Cloud SQL Instance
gcloud sql instances create griffin-dev-db \
    --database-version=MYSQL_8_0 \
    --edition=ENTERPRISE \
    --tier=db-custom-4-16384 \
    --storage-size=100GB \
    --region=$REGION \
    --root-password="my_secure_p@ssw0rd"

PROJECT_ID=$(gcloud config get-value project)
gcloud config set project $PROJECT_ID
gcloud auth login --no-launch-browser
gcloud sql connect griffin-dev-db --user=root --quiet

## Then execute this SQL script:
CREATE DATABASE wordpress;
CREATE USER "wp_user"@"%" IDENTIFIED BY "stormwind_rules";
GRANT ALL PRIVILEGES ON wordpress.* TO "wp_user"@"%";
FLUSH PRIVILEGES;


# Task5: Create Kubernetes cluster
CLUSTER_NAME="griffin-dev"
gcloud container clusters create $CLUSTER_NAME \
    --num-nodes=2 \
    --machine-type=e2-standard-4 \
    --zone=us-central1-c \
    --network=griffin-dev-vpc \
    --subnetwork=griffin-dev-wp

# Task6: Prepare the Kubernetes cluster
gcloud storage cp -r gs://spls/gsp321/wp-k8s .

gcloud iam service-accounts keys create key.json \
    --iam-account=cloud-sql-proxy@$GOOGLE_CLOUD_PROJECT.iam.gserviceaccount.com
kubectl create secret generic cloudsql-instance-credentials \
    --from-file key.json

# Task7: Create a WordPress deployment
cd wp-k8s
k apply -f wp-env.yaml
k apply -f wp-deployment.yaml
k apply -f wp-service.yaml

# Task8: Enable monitoring
gcloud beta monitoring uptime-checks create "WordPress development site Uptime Check" \
    --protocol=HTTP \
    --host=34.30.83.159 \
    --resource-type=uptime_url \
    --timeout=15s

# Task9
gcloud projects add-iam-policy-binding $GOOGLE_CLOUD_PROJECT \
    --member=user:student-03-09b585a15db9@qwiklabs.net \
    --role=roles/editor
