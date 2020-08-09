APPNAME := codecademy
DEV_IMAGE_NAME := codecademy-dev

.PHONY: build
build:
	@GO111MODULE=on go build -o $(APPNAME)

.PHONY: dev-image
dev-image:
	DOCKER_BUILDKIT=1 docker build -f dev.Dockerfile -t $(DEV_IMAGE_NAME) --ssh default .

.PHONY: clean
clean:
	rm coverage.*

.PHONY: fmt
fmt:
	go fmt `go list`/...

.PHONY: update-lint
update-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ci/ v1.27.0

.PHONY: lint
lint: fmt
	./ci/golangci-lint run -c ci/config.yml

.PHONY: tests
tests:
	@GO111MODULE=on go test ./... -v

.PHONY: test-coverage
test-coverage:
	go test -covermode=count -coverprofile=coverage.out -coverpkg=`go list ./... | grep -v '/scripts\|/protobuf' | tr "\n" ","` ./...
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html

.PHONY: mocks
mocks:
	@GO111MODULE=on mockery --dir=httpserver/controller --name=FileStore --inpackage
	@GO111MODULE=on mockery --dir=configuration --name=VarStore --inpackage
	@GO111MODULE=on mockery --dir=httpserver/controller --name=DBStore --inpackage

.PHONY: s3
s3:
	@aws --endpoint-url=http://localhost:4572 s3 ls s3://codecademy --recursive

.PHONY: db
db:
	@docker exec -i mysql mysql -uroot <<< "SELECT * FROM codecademy.pictures;"