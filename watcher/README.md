# Watcher for balancing endpoint

![](https://raw.githubusercontent.com/kelpi-io/kelpi.io/main/docs/watcher/structure.drawio.svg)

## TCP exmaple 

```Yaml
{
  "globalName": "gslb-operator.io",
  "balanceType": "wrr",
  "monitorType" : "tcp",
  "interval": 2,
  "monitor": {
    "interval": 2,
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
  "monitorType" : "http",
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
    "timeout": 2
  },
  "members": {
    "ip2" : {
      "ip": "104.196.232.237",
      "weight": 1
    }
  }
}
```

## Static check

```YAML
{
  "globalName": "gslb-operator.io",
  "balanceType": "wrr",
  "monitorType" : "static",
  "interval": 2,
  "monitor": {
    "enabled" : true
  },
  "members": {
    "ip2" : {
      "ip": "104.196.232.237",
      "weight": 1
    }
  }
}
```

## TODO
- [ ] Retries