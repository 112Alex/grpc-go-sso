# grpc-go-sso

**Статус:** проект в активной разработке
- JWT/refresh токены
- Документация по API и примеры клиентов
# grpc-go-sso

**Status:** In development

## Overview
gRPC-based authentication and authorization service for distributed systems and microservices.

## Technologies
- Go 1.24+
- gRPC ([grpc-go](https://github.com/grpc/grpc-go))
- Protobuf (`github.com/112Alex/protos`)
- slog
- Context7

## Structure
- `cmd/sso/main.go` — entry point
- `internal/app/` — app init, gRPC server
- `internal/services/auth/` — auth logic
- `internal/domain/models/` — domain models
- `config/` — configs

## Quick start
```sh
git clone https://github.com/112Alex/grpc-go-sso.git
cd grpc-go-sso
git checkout dev
go mod download
go run ./cmd/sso/main.go --config=config/local.yaml
```

## API
gRPC endpoints: login, register, isAdmin. Protobuf: `github.com/112Alex/protos`.

## Roadmap
- User storage
- JWT tokens
- Roles
- API docs

## References
- [gRPC-Go](https://github.com/grpc/grpc-go)
- Context7
import "google.golang.org/grpc"
