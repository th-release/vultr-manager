FROM golang:alpine
WORKDIR /app
COPY . .
RUN apk add --no-cache git
RUN go mod download
RUN go build -o vultr-manager
EXPOSE ${PORT:-3000}
CMD ["sh", "-c", "./vultr-manager"]