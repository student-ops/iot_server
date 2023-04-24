from(bucket: "vi_monitor")
  |> range(start: -1h)
  |> filter(fn: (r) => r._measurement == "vi_surroundings")
  |> filter(fn: (r) => r.user == "bar")
