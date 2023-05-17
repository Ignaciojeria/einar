FROM golang:latest

# Establece el directorio de trabajo
WORKDIR /go/src/app

# Copia los archivos del proyecto al contenedor
COPY . .

# Instala las dependencias del proyecto
RUN go get -d -v ./...

# Compila la aplicación
RUN go build -o /go/src/app/main .

# Expone el puerto 8080
EXPOSE 8080

# Define el comando para arrancar la aplicación
CMD ["/go/src/app/main"]