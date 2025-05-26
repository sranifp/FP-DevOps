FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod tidy

ENV DB_HOST=postgres
ENV DB_USER=pso
ENV DB_PASS=pso
ENV DB_NAME=pso
ENV DB_PORT=5432

CMD ["air"]