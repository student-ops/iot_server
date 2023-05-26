# azure

docker cp ./ <コンテナ名またはコンテナ ID>:/
docker run -it mcr.microsoft.com/azure-cli
docker cp ../vuoy_monitor_smaple 5478:/

親ディレクトリのファイルをコピー

//az cli をビルド
docker build -f azure/Dockerfile -t az-cli .

//az cli を dind と run
docker run --link dind:docker -v /var/run/docker.sock:/var/run/docker.sock -it az-cli

az commnads

```
az group create --name myResourceGroup --location eastus
az acr create --resource-group myResourceGroup --name my-first-acr  --sku Basic


az acr create --resource-group myResourceGroup --name maemuralaboacr  --sku Basic
```

docker tag myapp:latest maemuralab/myrepo:myapp
docker push maemuralab/myrepo:myapp

influxdb query

```
influx query --raw 'from(bucket:"vuoy_monitor") |> range(start:-1mo)'

```

```flux

influx query --raw 'from(bucket: "vuoy_monitor")
  |> range(start: -30d)
  |> filter(fn: (r) => r["_measurement"] == "vuoy_surroundings")
  |> filter(fn: (r) => r["_field"] == "AirPressure")
'
```

```flux
influx delete --org  iot --bucket vuoy_monitor --start '1970-01-01T00:00:00Z' --stop $(date +"%Y-%m-%dT%H:%M:%SZ")
```
