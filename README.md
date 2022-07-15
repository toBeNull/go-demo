## BOE
```shell
sh build.sh 

docker login --username=bobo1234@2000007798 cr-cn-beijing.volces.com
docker build -t cr-cn-beijing.volces.com/matthew-demo/go-demo:latest .
docker push cr-cn-beijing.volces.com/matthew-demo/go-demo:latest

kubectl delete -f manifest/deployment-orig.yml
kubectl apply -f manifest/deployment-orig.yml
```

## Online

```shell
sh build.sh 

docker login --username=pengjian@1085014 cr-cn-beijing.volces.com
docker build -t cr-cn-beijing.volces.com/mars-online/go-demo:latest .
docker push cr-cn-beijing.volces.com/mars-online/go-demo:latest

kubectl delete -f manifest/deployment-orig-online.yml
kubectl apply -f manifest/deployment-orig-online.yml
```


local:
```shell
docker run --name go-hello \
-p 8080:8080 \
-t \
--rm \
cr-cn-beijing.volces.com/matthew-demo/go-demo:latest
```