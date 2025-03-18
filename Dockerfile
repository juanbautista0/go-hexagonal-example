# Primera etapa: compilación
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Copiar y descargar dependencias primero para aprovechar el caché de Docker
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente y compilar el binario de la Lambda
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
RUN chmod +x bootstrap

# Segunda etapa: imagen final con el estándar de AWS Lambda
FROM public.ecr.aws/lambda/go:latest

# Copiar solo el binario compilado
COPY --from=builder /app/bootstrap /var/runtime/bootstrap

EXPOSE 8080

# AWS Lambda ejecuta automáticamente el binario llamado "bootstrap"
CMD [ "bootstrap" ]
