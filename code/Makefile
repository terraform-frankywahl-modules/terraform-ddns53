# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-s -w"

# AWS Lambda zip (Building the ZIP)
.PHONY: lambda
lambda:
	GOOS=linux GOARCH=amd64 go build -tags=lambda ${LDFLAGS} -o ./lambda cmd/lambda/*.go

# AWS Lambda zip (Building the ZIP in docker)
.PHONY: docker_lambda
docker_lambda:
	docker run -e GOOS=linux -e GOARCH=amd64 -v $$(pwd):/app -w /app golang:latest go build -tags=lambda $(LDFLAGS) -o lambda *.go

.PHONY: clean
clean:
	rm -rf lambda.zip lambda
