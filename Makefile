# Scanner
SONARQUBE_SCANNER_DOCKERFILE=".devcontainer/sonarqube-scanner.Dockerfile"
SONARQUBE_SCANNER_ENV_FILE=".devcontainer/sonarqube-scanner.env"
GO_COVERAGE_TEST_FILE="coverage.out"

# App Database
APP_DB_DOCKERFILE=".devcontainer/app-db.Dockerfile"
APP_DB_ENV_FILE=".devcontainer/app-db.env"

# App
POSTGRES_HOST=localhost
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_PORT=5432
APP_HOST=localhost
APP_PORT=3000


scanner-build:
	docker image build -t sonarqube-scanner:1.0 -f ${SONARQUBE_SCANNER_DOCKERFILE} .

scanner-run: scanner-build
	docker container run --env-file=${SONARQUBE_SCANNER_ENV_FILE} --rm sonarqube-scanner:1.0

app-db-build:
	docker image build -t app-db:1.0 -f ${APP_DB_DOCKERFILE} .

app-db-run: app-db-build
	docker container run --env-file=${APP_DB_ENV_FILE} -p "${POSTGRES_PORT}:${POSTGRES_PORT}" --rm app-db:1.0

go-test:
	go test -coverprofile=${GO_COVERAGE_TEST_FILE} -v ./...
	go tool cover -html=${GO_COVERAGE_TEST_FILE} -o coverage.html

go-run:
	\
		POSTGRES_HOST=${POSTGRES_HOST} \
		POSTGRES_USER=${POSTGRES_USER} \
		POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
		POSTGRES_PORT=${POSTGRES_PORT} \
		APP_HOST=${APP_HOST} \
		APP_PORT=${APP_PORT} \
		go run .
