# Global Server Load Balancer in Kubernetes

![](https://raw.githubusercontent.com/kelpi-io/kelpi.io/main/docs/components.drawio.svg)

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
  - ❌ Weight Round Robin
  - ❌ Failover group (active-backup)
  - ✅ Static (return all health endpoint)
  - ❌ Blank (return all endpoint)

- Fallback method
  - ❌ Return fallback endpoint
  - ❌ Return all endpoints
  - ❌ Refused query

- ❌ Automatic SOA serial
