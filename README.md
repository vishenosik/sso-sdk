# SSO-SDK

A package to create sso microservices clients. 
Clone this repo to start working.

## Taskfile

### Dependencies:

1. Swagger util [go-swagger](https://github.com/go-swagger/go-swagger) to generate golang REST client.
```bash
go install github.com/go-swagger/go-swagger/cmd/swagger@latest
```

2. Protobuf util [protoc](https://grpc.io/docs/protoc-installation/) to generate golang gRPC client.\

3. Setup .env file (lookup [example](sso-sdk/example.env)) either in root or sdk directory.
---

For better user-experience include Taskfile.yaml to your root Taskfile by adding following lines:
```yaml
includes:
  sdk:
    taskfile: ./sso-sdk/Taskfile.yaml
    optional: true
```
