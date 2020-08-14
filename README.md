# Sample Go AWS Lambda

Sample project to test AWS Lambda functions written in Go with [LocalStack](https://localstack.cloud)

## Install

- Clone this repository
- Install [Docker](https://docs.docker.com/engine/install/)
- Install [Docker Compose](https://docs.docker.com/compose/install/)
- Run `docker-compose up` to setup LocalStack

## Compile and deploy

- Compile and install with `make install`.
- If you do some changes, `make` should compile, build and deploy.

## Invoking function
```bash
aws --endpoint-url http://localhost:4566 lambda invoke \
  --function-name sample-lambda \
  --payload '{"firstname": "foo", "lastname": "bar", "age": 41}' \
  output
```
Should produce an output with contents:
```
"{\"name\":\"foo bar\",\"age\":41}"
```