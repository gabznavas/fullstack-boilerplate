# Estágio de construção usando Go Alpine
FROM golang:1.24.6-alpine AS builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos de dependências do Go
COPY go.mod go.sum ./

# Baixa as dependências do módulo Go
RUN go mod download

# Copia todo o código fonte para o container
COPY . .

# Compila a aplicação Go com otimizações para produção
RUN go build -o main main.go

# Estágio final usando imagem vazia para container mínimo
FROM scratch

# Copia o binário compilado do estágio de construção para a imagem final
COPY --from=builder /app/main /main

# Define o comando de entrada do container
CMD ["/main"]

# Para subir a imagem aws-api para o Docker Hub, execute os seguintes comandos:
# 1. Faça o build da imagem com tag: docker build --tag gabrielnavas/aws-api:v1.0.0 .
# 2. Faça login no Docker Hub: docker login
# 3. Suba a imagem: docker push gabrielnavas/aws-api:v1.0.0
# 5. Suba a tag latest: docker push seu-usuario/aws-api:latest


# Para executar o container na porta 8080:
# docker run -p 8080:8080 gabrielnavas/aws-api:v1.0.0