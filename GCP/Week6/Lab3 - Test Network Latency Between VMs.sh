gcloud compute instances create us-test-01 \
--subnet subnet-us-west1 \
--zone us-west1-a \
--machine-type e2-standard-2 \
--tags ssh,http,rules

gcloud compute instances create us-test-02 \
--subnet subnet-europe-west1 \
--zone europe-west1-d \
--machine-type e2-standard-2 \
--tags ssh,http,rules

gcloud compute instances create us-test-03 \
--subnet subnet-us-east4 \
--zone us-east4-a \
--machine-type e2-standard-2 \
--tags ssh,http,rules

gcloud compute instances create us-test-04 \
--subnet subnet-us-west1 \
--zone us-west1-b \
--machine-type e2-standard-2 \
--tags ssh,http,rules

# Inside VMs

## Using FQDN
ping -c 3 us-test-02.subnet-europe-west1

## Install traceroute
sudo apt-get update
sudo apt-get -y install traceroute mtr tcpdump iperf whois host dnsutils siege

## Trace
traceroute us-test-02.subnet-europe-west1

## Use iperf to test performance
sudo apt-get update
sudo apt-get -y install traceroute mtr tcpdump iperf whois host dnsutils siege

## Run server
iperf -s #run in server mode
iperf -c us-test-01.us-west1-a #run in client mode

iperf -s -u #iperf server side with UDP
iperf -c us-test-01.us-west1-a -u -b 2G #iperf client side - send 2 Gbits/s
iperf -c us-test-01.us-west1-a -P 20 #iperf client side - send 20 in parallel
