## 開発中

dev/slack ブランチ

docker run -d -p 3000:3000 grafana/grafana-oss:main

## Getting started

import grafana volume on

- ./grafana/grafana_volume:/var/lib/grafansa

let's compose up from local file

```
docker compose -f docker-compose.local.yml up -d

cd server_tests/sample_generator/

python3 sample_generator.py
```

then open http://localhost:3000

## INfluxdb query

```

from(bucket: "vuoy_monitor")
|> range(start: v.timeRangeStart, stop: v.timeRangeStop)
|> filter(fn: (r) => r["_measurement"] == "vuoy_surroundings")
|> filter(fn: (r) => r["_field"] == "Rssi")
|> yield(name: "mean")

```

k6 commnad

```

k6 run --vus 10 --duration 30s ./server_tests/sample_generator/sample_generator.js

```
