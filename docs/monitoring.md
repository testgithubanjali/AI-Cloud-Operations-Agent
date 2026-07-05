# Monitoring

The cluster is monitored using Prometheus.

Dashboards are available in Grafana.

CPU and Memory metrics come from Metrics Server.

Alerts are managed by Alertmanager.

Useful commands:

```bash
kubectl top pods
kubectl top nodes
```