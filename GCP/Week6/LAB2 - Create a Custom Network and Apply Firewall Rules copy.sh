VPC_NAME=taw-custom-network
gcloud compute networks create $VPC_NAME --subnet-mode=custom

SUBNET1_REGION="us-central1"
SUBNET1_NAME="subnet-$SUBNET1_REGION"
SUBNET1_RANGE="10.0.0.0/16"
gcloud compute networks subnets create $SUBNET1_NAME --network=$VPC_NAME --region=$SUBNET1_REGION --range=$SUBNET1_RANGE

SUBNET2_REGION="us-east4"
SUBNET2_NAME="subnet-$SUBNET2_REGION"
SUBNET2_RANGE="10.1.0.0/16"
gcloud compute networks subnets create $SUBNET2_NAME --network=$VPC_NAME --region=$SUBNET2_REGION --range=$SUBNET2_RANGE

SUBNET3_REGION="us-west1"
SUBNET3_NAME="subnet-$SUBNET3_REGION"
SUBNET3_RANGE="10.2.0.0/16"
gcloud compute networks subnets create $SUBNET3_NAME --network=$VPC_NAME --region=$SUBNET3_REGION --range=$SUBNET3_RANGE

gcloud compute firewall-rules create "nw101-allow-http" --direction=INGRESS --priority=1000 --network=$VPC_NAME --action=ALLOW --rules=tcp:80 --source-ranges=0.0.0.0/0 --target-tags=http

gcloud compute firewall-rules create "nw101-allow-icmp" --direction=INGRESS --priority=1000 --network=$VPC_NAME --action=ALLOW --rules=icmp --source-ranges=0.0.0.0/0 --target-tags=rules

gcloud compute firewall-rules create "nw101-allow-ssh" --direction=INGRESS --priority=1000 --network=$VPC_NAME --action=ALLOW --rules=tcp:22 --source-ranges=0.0.0.0/0 --target-tags=ssh

gcloud compute firewall-rules create "nw101-allow-rdp" --direction=INGRESS --priority=1000 --network=$VPC_NAME --action=ALLOW --rules=tcp:3389 --source-ranges=0.0.0.0/0 --target-tags=rdp

gcloud compute firewall-rules create "nw101-allow-internal" --direction=INGRESS --priority=1000 --network=$VPC_NAME --action=ALLOW --rules=tcp:0-65535,udp:0-65535,icmp --source-ranges=10.0.0.0/16,10.1.0.0/16,10.2.0.0/16