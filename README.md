### prom-remote-writer

Write a metric to a remote Prometheus server

### Usage

Add environment vars:

```
PROM_URL=https://<prometheus hostname>/api/v1/write
PROM_USERNAME=<basic auth username>
PROM_PASSWORD=<basic auth password>
```

Then run:

```
./prom-remote-writer --metric_name example_metric_name --metric_labels "version=1.0.0" --value 1
```
