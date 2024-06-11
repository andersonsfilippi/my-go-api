#imagem
FROM golang:1.16

#diretorio de trabalho
WORKDIR /app

#copiar os arquivos go.mod e go.sum e baixar as dependências
COPY go.mod go.sum ./
RUN go mod download

#copiar o código fonte para o container
COPY . .

#Compilar a aplicação
RUN go build -o main .

#definir porta que será exposta
EXPOSE 8080

#comando para rodar a aplicação
CMD [ "./main" ]