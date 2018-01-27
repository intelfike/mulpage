go build
sudo docker build -t asia.gcr.io/intelfike-428ac/mulpage .
sudo /usr/local/google-cloud-sdk/bin/gcloud docker -- push asia.gcr.io/intelfike-428ac/mulpage
kubectl delete pod --all
kubectl run mulpage --image=asia.gcr.io/intelfike-428ac/mulpage --port=80
kubectl expose deployment mulpage --port 80 --type LoadBalancer --load-balancer-ip="35.201.227.241"
