# То что будет установенно в контейнер [ЯП:VERSION]
FROM golang:1.21
# Устанавливаю рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum в рабочую директорию
COPY go.* ./

# Скачиваем зависимости
RUN go mod download

COPY . .

# Компилируем приложение для продакшена
RUN go build -o ascii-art ./cmd/

# Определите порт, на котором будет работать приложение
EXPOSE 8080

# Запустите скомпилированное приложение
CMD ["./ascii-art"]

# Одна из разновидностей метаданных. Конкретно эти называются "постоянными"
LABEL maintainer="Daniil Mjodov <GitTea - dpruel>"
LABEL description="ASCII Art Web Application"
