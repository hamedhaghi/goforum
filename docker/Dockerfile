FROM golang:alpine3.19

LABEL author="Hamed Haghi" email="hamedhaghi07@gmail.com"

RUN apk update && apk upgrade

RUN apk --no-cache add curl

RUN go install github.com/pilu/fresh@latest

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2

WORKDIR /app

COPY . .

EXPOSE 3000

CMD ["fresh"]
