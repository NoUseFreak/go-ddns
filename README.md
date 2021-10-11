# go-ddns

`go-ddns` provides a way to update your DNS record.

## How it works

If will resolve the public ip by trying to reach the internet.
It will check if the dns record needs updating and update if needed.

```
timeout: 60
spec:
- type: aws-route53
  domain: office.example.com
  options:
    AWS_ACCESS_KEY_ID: xxxxxxxxxxxxxxxxxxxx
    AWS_SECRET_ACCESS_KEY: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
    zoneId: xxxxxxxxxxxxx
```

```
go-ddns config.yml
```

```
docker run -d -v $(pwd)/config.yml:/config.yml ghcr.io/nousefreak/go-ddns
```

## Supported adapters

Currently only `aws-route53` is supported.


