# Global Server Load Balancer in Kubernetes

![](https://raw.githubusercontent.com/vaishutin/gslb-operator/main/docs/components.drawio.svg)

## Features
- Support 3 type of HealthCheck
  - TCP - tcp query on port
  - HTTP - http query with support TLS, path, and set custorm headers
  - Static - all endpoint permanent enabled or disabled
