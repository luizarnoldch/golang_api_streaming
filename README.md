# Título de tu Proyecto

Contenido...

## Estadísticas del Proyecto

| Métrica       | Valor         |
|---------------|---------------|
| Cobertura     | 76.87%        |
| Otra Métrica  | Otro Valor    |


## Requirements

- Make
- Go 20+
- Docker
- AWS CLI v2

## Dependencies

```shell

make init
make update
go install github.com/vektra/mockery/v2@v2.39.2
make mock

```


## Run Code

```shell

make dynamo-up      # Run DynamoDB locally
make unit           # Run unit testing
make coverage       # Update the coverage code
make deploy         # Deploy on AWS

```

## TODO

- Separate Unit testing with Integration Testing
- Make End-To-End testing with API-Gateway Locally or Lambda Locally