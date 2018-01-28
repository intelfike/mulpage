# バイナリ
go build

# バージョン管理
version=$(echo `cat version.txt`+0.1 | bc)
echo $version > version.txt
sed -i -e "s/mulpage:[0-9]\.[0-9]*/mulpage:$version/" deployment.yaml
# コンテナ
sudo docker build -t asia.gcr.io/intelfike-428ac/mulpage:$version .
sudo /usr/local/google-cloud-sdk/bin/gcloud docker -- push asia.gcr.io/intelfike-428ac/mulpage:$version
kubectl apply -f deployment.yaml

# ↓使えそうだったけどダメだったコマンド
# kubectl rolling-update mulpage -f=mulpage-pod.yaml
# kubectl delete pod --all
# kubectl run mulpage --image=asia.gcr.io/intelfike-428ac/mulpage --port=80
# kubectl expose deployment mulpage --port 80 --type LoadBalancer --load-balancer-ip="35.201.227.241"
