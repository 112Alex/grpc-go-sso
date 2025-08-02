# grpc-go-sso

**Статус:** проект в активной разработке

## Описание

`grpc-go-sso` — это сервис авторизации, реализованный на Go с использованием gRPC. Проект предназначен для централизованной аутентификации и авторизации пользователей в распределённых системах и микросервисной архитектуре.

В текущей версии реализована базовая структура приложения, запуск gRPC-сервера и заготовки для Auth API (login, register, isAdmin). В дальнейшем планируется расширение функциональности, интеграция с хранилищем пользователей, поддержка JWT и ролей.

## Основные технологии
- Go 1.24+
- gRPC ([grpc-go](https://github.com/grpc/grpc-go))
- Протоколы описываются через protobuf (см. зависимость `github.com/112Alex/protos`)
- Логирование через slog

## Структура проекта
- `cmd/sso/main.go` — точка входа, запуск сервера, graceful shutdown
- `internal/app/` — инициализация приложения, запуск gRPC
- `internal/config/` — работа с конфигами
- `internal/config/grpc/auth/` — реализация gRPC Auth API
- `lib/logger/` — обработчики логирования
- `config/` — конфигурационные файлы (пример: `local.yaml`)

## Быстрый старт (dev)
1. Установите Go 1.24+
2. Клонируйте репозиторий:
   ```sh
   git clone https://github.com/112Alex/grpc-go-sso.git
   cd grpc-go-sso
   git checkout dev
   ```
3. Установите зависимости:
   ```sh
   go mod download
   ```
4. Проверьте/отредактируйте конфиг `config/local.yaml` при необходимости.
5. Запустите сервис:
   ```sh
   go run ./cmd/sso/main.go --config=config/local.yaml
   ```

## gRPC API
Сервис реализует gRPC-интерфейс авторизации (login, register, isAdmin). Прототипы методов и сообщения описаны в отдельном репозитории протоколов (`github.com/112Alex/protos`).

Пример запуска gRPC сервера:
```go
import "google.golang.org/grpc"

s := grpc.NewServer()
// регистрация сервисов
go s.Serve(lis)
```

## Планы
- Реализация полноценного хранилища пользователей
- JWT/refresh токены
- Ролевое управление
- Документация по API и примеры клиентов

## Контакты
- Вопросы и предложения: issues в GitHub

---

> Основано на best practices из официальной документации [gRPC-Go](https://github.com/grpc/grpc-go) и Context7.
