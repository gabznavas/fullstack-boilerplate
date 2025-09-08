# Etapa 1: build
FROM node:20-alpine AS builder
WORKDIR /app

# Permite passar a variável no build
ARG NEXT_PUBLIC_API_URL
ENV NEXT_PUBLIC_API_URL=$NEXT_PUBLIC_API_URL

COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Etapa 2: produção
FROM node:20-alpine
WORKDIR /app

# Mantém a variável em produção (opcional, mas documenta)
ARG NEXT_PUBLIC_API_URL
ENV NEXT_PUBLIC_API_URL=$NEXT_PUBLIC_API_URL

COPY --from=builder /app/package*.json ./
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public

# Expõe a porta 3000 (documentação / referência)
EXPOSE 3000

CMD ["npm", "run", "start"]

# Para buildar a imagem passando a variável de ambiente:
# docker build --build-arg NEXT_PUBLIC_API_URL=http://localhost:8080 -t gabrielnavas/aws-web:v1.0.0 .

# Para subir a imagem AWS Web para o Docker Hub:
# docker push gabrielnavas/aws-web:v1.0.0
# docker tag gabrielnavas/aws-web:v1.0.0 gabrielnavas/aws-web:latest
# docker push gabrielnavas/aws-web:latest

# Para executar o container na porta 3000:
# docker run -p 3000:3000 gabrielnavas/aws-web:v1.0.0
