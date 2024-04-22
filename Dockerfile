FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . /app/

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o go-boilerplate main.go

FROM golang:1.20-alpine

COPY --from=builder ./app/ /usr/local/bin

# Pass flag with environment variable
ENTRYPOINT ["go-boilerplate", "-mode=production"]  
