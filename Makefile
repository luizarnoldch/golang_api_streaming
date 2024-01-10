TEMPLATE_FILE := deployment/cloudformation/main.yml
STACK_NAME := streaming-api

DYNAMO-LOCAL := deployment/docker/dynamo-local.yml

init:
	go mod init main
update:
	go mod tidy
mock:
	mockery --all --output ./mocks
coverage:
	go test ./... -cover
cover-unit:
	go test ./... -coverprofile=coverage-report.out
	go tool cover -html=coverage-report.out -o coverage-report.html
	go tool cover -func=coverage-report.out
	go run ./utils/coverage/coverage_generator.go
build:
	chmod +x -R ./scripts/
	./scripts/build.sh
unit:
	go test ./...
f_test:
	chmod +x -R ./scripts/
	./scripts/func_test.sh
deploy:
	sam deploy --template-file $(TEMPLATE_FILE) --stack-name $(STACK_NAME) --capabilities CAPABILITY_NAMED_IAM --resolve-s3
destroy:
	aws cloudformation delete-stack --stack-name $(STACK_NAME)
dynamo-up:
	docker-compose -f ${DYNAMO-LOCAL} up -d
dynamo-stop:
	docker-compose -f ${DYNAMO-LOCAL} stop
dynamo-start:
	docker-compose -f ${DYNAMO-LOCAL} start
dynamo-destroy:
	make dynamo-stop
	docker-compose -f ${DYNAMO-LOCAL} down -v && docker rmi $$(docker images -q amazon/dynamodb-local) || true
e2e:
	make unit
	make coverage
	make build
	make deploy
	sleep 5s
	make f_test
