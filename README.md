# Prometheus F6005 (v3)

A theoretically simple way to export metrics from your ZTE F6005v3 ONT to Prometheus.

## Required environment variables

| Name       | Description             | Default Value      |
|------------|-------------------------|--------------------|
| `ENDPOINT` | HTTP address to the ONT | http://192.168.1.1 |
| `USERNAME` | Username for the ONT    | admin              |
| `PASSWORD` | Password for the ONT    | admin              |

## Usage

Example docker-compose section:

```yaml
  f6005_exporter:
    restart: always
    image: ghcr.io/lucathehacker/prometheusf6005
    environment:
      - ENDPOINT=http://192.168.1.1
      - ONT_USERNAME=admin
      - ONT_PASSWORD=admin
    expose:
      - 80
```

###### The code sucks, I know. I wrote it with a bad headache and no sleep.
