# Task 1
REGION="us-central1"
ZONE="us-central1-c"
gcloud config set compute/region $REGION
gcloud config set compute/zone $ZONE

CLUSTER_NAME="onlineboutique-cluster-198"
gcloud container clusters create $CLUSTER_NAME --zone=$ZONE --machine-type=e2-standard-2 --num-nodes=2

kubectl create namespace dev && kubectl create namespace prod

git clone https://github.com/GoogleCloudPlatform/microservices-demo.git && \
cd microservices-demo && kubectl apply -f ./release/kubernetes-manifests.yaml --namespace dev

kubectl get service frontend-external -o wide  --namespace dev

# Task 2
NODE_POOL_NAME="optimized-pool-4656"
gcloud container node-pools create $NODE_POOL_NAME --cluster $CLUSTER_NAME --zone $ZONE --machine-type=custom-2-3584 --num-nodes=2

for node in $(kubectl get nodes -l cloud.google.com/gke-nodepool=default-pool -o=name); do
  kubectl cordon "$node";
done

kubectl create poddisruptionbudget kube-dns-pdb --namespace=kube-system --selector k8s-app=kube-dns --max-unavailable 1
kubectl create poddisruptionbudget prometheus-pdb --namespace=kube-system --selector k8s-app=prometheus-to-sd --max-unavailable 1
kubectl create poddisruptionbudget kube-proxy-pdb --namespace=kube-system --selector component=kube-proxy --max-unavailable 1
kubectl create poddisruptionbudget metrics-agent-pdb --namespace=kube-system --selector k8s-app=gke-metrics-agent --max-unavailable 1
kubectl create poddisruptionbudget metrics-server-pdb --namespace=kube-system --selector k8s-app=metrics-server --max-unavailable 1
kubectl create poddisruptionbudget fluentd-pdb --namespace=kube-system --selector k8s-app=fluentd-gke --max-unavailable 1
kubectl create poddisruptionbudget backend-pdb --namespace=kube-system --selector k8s-app=glbc --max-unavailable 1
kubectl create poddisruptionbudget kube-dns-autoscaler-pdb --namespace=kube-system --selector k8s-app=kube-dns-autoscaler --max-unavailable 1
kubectl create poddisruptionbudget stackdriver-pdb --namespace=kube-system --selector app=stackdriver-metadata-agent --max-unavailable 1
kubectl create poddisruptionbudget event-pdb --namespace=kube-system --selector k8s-app=event-exporter --max-unavailable 1

kubectl create poddisruptionbudget event-pdb --namespace=gmp-system --selector app.kubernetes.io/name=gmp-operator --max-unavailable 1

for node in $(kubectl get nodes -l cloud.google.com/gke-nodepool=default-pool -o=name); do
  kubectl drain --force --ignore-daemonsets --grace-period=10 "$node" --delete-emptydir-data;
done

gcloud container node-pools delete default-pool --cluster $CLUSTER_NAME

# Task 3
kubectl create poddisruptionbudget onlineboutique-frontend-pdb --selector app=frontend --min-available 1 --namespace dev
kubectl get poddisruptionbudget --namespace dev


# Open editor
# Find:
#   image: us-central1-docker.pkg.dev/google-samples/microservices-demo/frontend:v0.10.3
# Replace it with:
#   image: gcr.io/qwiklabs-resources/onlineboutique-frontend:v2.1
#   ImagePullPolicy: Always
kubectl apply -f ./release/kubernetes-manifests.yaml --namespace dev

# Task 4
MAX_REPLICAS=12
kubectl autoscale deployment frontend --cpu-percent=50 --min=1 --max=$MAX_REPLICAS --namespace dev

gcloud beta container clusters update $CLUSTER_NAME --zone $ZONE --enable-autoscaling --min-nodes 1 --max-nodes 6

## Exec load test
kubectl get service frontend-external -o wide --namespace dev
YOUR_FRONTEND_EXTERNAL_IP=34.27.201.250
kubectl exec $(kubectl get pod --namespace=dev | grep 'loadgenerator' | cut -f1 -d ' ') -it --namespace=dev -- bash -c 'export USERS=8000; locust --host="http://34.27.201.250" --headless -u "8000" 2>&1'


kubectl autoscale deployment recommendationservice --cpu-percent=50 --min=1 --max=5 --namespace dev

# Task 5
gcloud container clusters update $CLUSTER_NAME --zone $ZONE \
    --enable-autoprovisioning \
    --min-cpu 1 \
    --min-memory 2 \
    --max-cpu 45 \
    --max-memory 160