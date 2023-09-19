# Chameleon

A DDNS client in Go.

> [!NOTE]
> Only Google Domains is supported.

## Usage

### Docker

#### Build

```
make docker
```

#### Run

```
docker run -d chameleon:latest \
  -u <DDNS username> \
  -p <DDNSpassword> \
  -h <hostname to update>
```
