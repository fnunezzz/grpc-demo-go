# grpc-demo-go

This is a simple grpc demo project made in go with JWT authentication.

## Instalation

```bash
go mod tidy
```

You will also need to import the proto files in your application of choice. I personally prefer [Postman](https://www.postman.com/) and [you can follow postman tutorial here](https://learning.postman.com/docs/sending-requests/grpc/grpc-request-interface/). These proto files are located at `internal/proto/build`.

## Running

```bash
# Start a local postgres instance and run migrations located at cmd/migration
docker-compose up
```

```bash
# Start the go API
make dev
```

You should now be able to start making API calls.

-   create a user with `SignUp` method.
-   login with that user using `SignIn` method
-   use an auth only route using `HealthCheck` method

### SignUp

```json
{
    "email": "email@email",
    "password": "abc12345678",
    "username": "grpc-demo"
}
```

### SignIn

```json
{
    "password": "abc12345678",
    "username": "filipe.nunez"
}
```

### HealthCheck

Header: `Authorization: Bearer ${token}`

## Test

```bash
make test
```

## Generate pb.go files

This is a basic command line argument to generate `pb.go` files.

```bash
protoc --proto_path=internal/proto/build --go_out=internal/proto/gen --go_opt=paths=source_relative --go-grpc_out=internal/proto/gen --go-grpc_opt=paths=source_relative internal/proto/build/p_user.proto;
```

```bash
protoc --proto_path=internal/proto/build --go_out=internal/proto/gen --go_opt=paths=source_relative --go-grpc_out=internal/proto/gen --go-grpc_opt=paths=source_relative internal/proto/build/p_app.proto;
```
