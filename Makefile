SONARQUBE_SCANNER_DOCKERFILE=".devcontainer/sonarqube-scanner.Dockerfile"
GO_COVERAGE_TEST_FILE="coverage.out"
ENV_FILE=".env"


docker-build:
	docker image build -t sonarqube-scanner:1.0 -f ${SONARQUBE_SCANNER_DOCKERFILE} .

docker-run: docker-build
	docker container run --env-file=${ENV_FILE} --rm sonarqube-scanner:1.0

go-test:
	go test -coverprofile=${GO_COVERAGE_TEST_FILE} -v ./...
	go tool cover -html=${GO_COVERAGE_TEST_FILE} -o coverage.html
