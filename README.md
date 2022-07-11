```shell
sh build.sh 

docker login --username=bobo1234@2000007798 cr-cn-beijing.volces.com
docker build -t cr-cn-beijing.volces.com/matthew-demo/go-demo:latest .
docker push cr-cn-beijing.volces.com/matthew-demo/go-demo:latest
```


local:
```shell
docker run --name go-hello \
-p 8080:8080 \
-t \
--rm \
cr-cn-beijing.volces.com/matthew-demo/go-demo:latest
```