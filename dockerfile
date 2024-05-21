FROM golang:1.19-alpine

WORKDIR /app

# Copie o arquivo go.mod e go.sum e baixe as dependências
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copie o código fonte
COPY . .

# Copie o arquivo .env
COPY .env .env

# Construa a aplicação
RUN go build -o main .

# Execute a aplicação
CMD ["./main"]
