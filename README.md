# SSO-SDK

A package to create sso microservices clients.\
Clone this repo to start working.

## Taskfile

### Dependencies

#### Swagger util [oapi-codegen](github.com/oapi-codegen/oapi-codegen) to generate golang REST client

```bash
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest (v2.5.0)
```

#### Protobuf util [protoc](https://grpc.io/docs/protoc-installation/) to generate golang gRPC client

plugins:

```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

#### Setup .env file (lookup [example](sso-sdk/example.env)) either in root or sdk directory

---


