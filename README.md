# Global Server Load Balancer in Kubernetes

![](https://raw.githubusercontent.com/vaishutin/gslb-operator/main/docs/components.drawio.svg)

## Features
- 👨‍💻 Working as PowerDNS remote backend
- ❌ Balancing pools as kubernetes manifests
- ❌ Prometheus metrics exporter
- ❌ Api for managing pools (manage over k8s manifests)
- ❌ UI for manipulation pools and reports
- Support 3 type of HealthCheck
  - ✅ TCP - tcp query on port
  - ✅ HTTP - http query with support TLS, path, and set custorm headers
  - ✅ Static - all endpoint permanent enabled or disabled

- LoadBalancing Methods
  - 👨‍💻 Static (return all health endpoint)
  - ❌ Weight Round Robin
  - ❌ Failover group (active-backup)

- Fallback method
  - ❌ Return fallback endpoint
  - ❌ Return all endpoints
  - ❌ Refused query

- ❌ Automatic SOA serial
