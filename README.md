## 開発中
dev/slack ブランチ

docker run -d -p 3000:3000 grafana/grafana-oss:main

## INfluxdb query

```
from(bucket: "vuoy_monitor")
  |> range(start: v.timeRangeStart, stop: v.timeRangeStop)
  |> filter(fn: (r) => r["_measurement"] == "vuoy_surroundings")
  |> filter(fn: (r) => r["_field"] == "Rssi")
  |> yield(name: "mean")

  ```