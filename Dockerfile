FROM golang:alpine
WORKDIR /app
COPY . .
RUN apk add --no-cache git
RUN go mod download
RUN go build -o vultr-manager
EXPOSE ${PORT:-8080}
CMD ["sh", "-c", "./vultr-manager --API ${API} --PASSWORD ${PASSWORD} --PORT ${PORT:-8080} --DATABASE_ADDR=${DATABASE_ADDR} --DATABASE_USER=${DATABASE_USER} --DATABASE_PASSWORD=${DATABASE_PASSWORD} --DATABASE_SCHEMA=${DATABASE_SCHEMA}"]