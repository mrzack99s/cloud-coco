# API

### Pre-requisite
1. Go Version 1.17.x

### How to use

#### How to install package
```shell
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
go install ./coco
```

#### How to generate-api-docs
```shell
swag init --pd --parseDepth 1 -g ./coco/launch.go
```

#### How to generate-ts
```shell
npx swagger-typescript-api -p ./docs/swagger.json -o <ui destination>  --axios
```

#### How to serve
```shell
go run ./coco
```