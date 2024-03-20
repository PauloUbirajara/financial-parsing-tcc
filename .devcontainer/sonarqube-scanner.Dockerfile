# Build
FROM golang:1.21.8 AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .
RUN go test -coverprofile=coverage.out -v ./...
RUN go build -o main .

# Run SonarQube scanner and generate score
FROM sonarsource/sonar-scanner-cli:5.0

COPY --from=builder /app .
CMD sonar-scanner \
  -X \
  -Dsonar.projectKey="$SONARQUBE_SCANNER_PROJECT_KEY" \
  -Dsonar.sources="$SONARQUBE_SCANNER_SOURCES" \
  -Dsonar.login="$SONARQUBE_SCANNER_LOGIN" \
  -Dsonar.host.url="$SONARQUBE_SCANNER_HOST"
