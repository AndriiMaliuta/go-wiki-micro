# syntax=docker/dockerfile:1
FROM golang:1.17-alpine
WORKDIR /app

COPY . .
#COPY go.mod ./
#COPY go.sum ./
RUN go mod download

#COPY *.go ./
RUN go build -o /go-app

EXPOSE 4000

CMD [ "/go-app" ]