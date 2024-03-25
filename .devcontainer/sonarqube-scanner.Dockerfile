# Build
FROM golang:1.21.8 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go test -coverprofile=coverage.out -v ./...
RUN go build -o main .

# Run SonarQube scanner and generate score
FROM sonarsource/sonar-scanner-cli:5.0

COPY --from=builder /app /usr/src

CMD sonar-scanner \
  -Dsonar.projectKey=financial-parsing \
  -Dsonar.sources=. \
  -Dsonar.host.url="${SONARQUBE_SCANNER_HOST}" \
  -Dsonar.login="${SONARQUBE_SCANNER_LOGIN}"
