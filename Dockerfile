# build stage
# get golang 
FROM golang:1.18-alpine AS builder
# working directory 
WORKDIR /app
# copy all files to work directory
COPY . .
# build executables file
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz



# run stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY wait-for.sh .
COPY start.sh .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY app.env .
COPY DB/migration ./migration

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]