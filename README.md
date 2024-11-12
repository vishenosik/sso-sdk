# SSO-SDK

A package to create sso microservices clients.\
Clone this repo to start working.

## Taskfile

### Dependencies

#### Swagger util [go-swagger](https://github.com/go-swagger/go-swagger) to generate golang REST client

```bash
go install github.com/go-swagger/go-swagger/cmd/swagger@latest
```

#### Protobuf util [protoc](https://grpc.io/docs/protoc-installation/) to generate golang gRPC client

plugins:

```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

#### Setup .env file (lookup [example](sso-sdk/example.env)) either in root or sdk directory

---

For better user-experience include Taskfile.yaml to your root Taskfile by adding following lines:

```yaml
includes:
  sdk:
    taskfile: ./sso-sdk/Taskfile.yaml
    optional: true
```
