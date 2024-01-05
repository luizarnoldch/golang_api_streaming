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
dynamo-local:
	docker-compose -f ${DYNAMO-LOCAL} up
e2e:
	make unit
	make coverage
	make build
	make deploy
	sleep 5s
	make f_test
