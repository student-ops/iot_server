# azure 
docker cp ./  <コンテナ名またはコンテナID>:/
docker run -it mcr.microsoft.com/azure-cli
docker cp ../vuoy_monitor_smaple  5478:/


親ディレクトリのファイルをコピー


//az cliをビルド
docker build -f azure/Dockerfile -t az-cli .

//az cliをdindとrun
docker run --link dind:docker -v /var/run/docker.sock:/var/run/docker.sock -it az-cli



az commnads

```
az group create --name myResourceGroup --location eastus
az acr create --resource-group myResourceGroup --name my-first-acr  --sku Basic


az acr create --resource-group myResourceGroup --name maemuralaboacr  --sku Basic
```


docker tag myapp:latest maemuralab/myrepo:myapp
docker push maemuralab/myrepo:myapp
