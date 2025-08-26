# Task 1. Create networks
VPC_NAME=vpc-network-lrx9
gcloud compute networks create $VPC_NAME --subnet-mode=custom

SUBNET1_REGION="us-central1"
SUBNET1_NAME="subnet-a-y1yz"
SUBNET1_RANGE="10.10.10.0/24"
gcloud compute networks subnets create $SUBNET1_NAME --network=$VPC_NAME --region=$SUBNET1_REGION --range=$SUBNET1_RANGE

SUBNET2_REGION="europe-west4"
SUBNET2_NAME="subnet-b-65gx"
SUBNET2_RANGE="10.10.20.0/24"
gcloud compute networks subnets create $SUBNET2_NAME --network=$VPC_NAME --region=$SUBNET2_REGION --range=$SUBNET2_RANGE

# Task 2. Add firewall rules
gcloud compute firewall-rules create "sfgp-firewall-ssh" --direction=INGRESS --priority=1000 \
    --network=$VPC_NAME --action=ALLOW --rules=tcp:22 --source-ranges=0.0.0.0/0

gcloud compute firewall-rules create "ltjb-firewall-rdp" --direction=INGRESS --priority=65535 \
    --network=$VPC_NAME --action=ALLOW --rules=tcp:3389 --source-ranges=0.0.0.0/24

gcloud compute firewall-rules create "zusq-firewall-icmp" --direction=INGRESS --priority=1000 \
    --network=$VPC_NAME --action=ALLOW --rules=icmp --source-ranges=10.10.10.0/24,10.10.20.0/24

# Task 3. Add VMs to your network
gcloud compute instances create  us-test-01 \
    --subnet $SUBNET1_NAME \
    --zone us-central1-a \
    --machine-type e2-standard-2

gcloud compute instances create  us-test-02 \
    --subnet $SUBNET2_NAME \
    --zone europe-west4-a \
    --machine-type e2-standard-2

gcloud compute instances list
