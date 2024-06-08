# Global Server Load Balancer in Kubernetes

![](https://raw.githubusercontent.com/vaishutin/gslb-operator/main/docs/img/components.png)

## TCP exmaple 

```Yaml
{
  "globalName": "gslb-operator.io",
  "balanceType": "wrr",
  "type" : "tcp",
  "interval": 2,
  "monitor": {
    "interval": 2,
    "retries": 2,
    "port": 80
  },
  "members": {
    "ip1" : {
      "ip": "1.1.1.1",
      "weight": 1
    },
    "ip2" : {
      "ip": "1.1.1.2",
      "weight": 1
    }
  }
}
```

## HTTP check

```YAML
{
  "globalName": "gslb-operator.io",
  "balanceType": "wrr",
  "type" : "http",
  "interval": 2,
  "monitor": {
    "https" : true,
    "host" : "dummyjson.com",
    "path" : "/test",
    "headers" : {
        "My-token" : "token"
        "User-Agnet" : "gslb-operatro"
    },
    "port": 443,
    "valid_codes": [200, 201],
    "timeout": 2,
    "retries": 2
  },
  "members": {
    "ip2" : {
      "ip": "104.196.232.237",
      "weight": 1
    }
  }
}
```
