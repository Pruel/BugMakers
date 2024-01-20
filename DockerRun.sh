#!/bin/bash

# Называем образ
Data_Image="ascii-art-app"

# Команда docker build собирает образ Docker из Dockerfile
# флаг -t присваивает этому образу имя $Data_Image
docker build -t $Data_Image .

# Запускаем контейнер из образа
# флаг -d запускает контейнер в фоновом режиме. 
# флаг -p 8080:8080 перенаправляет порт 8080 хоста на порт 8080 контейнера
docker run -d -p 8080:8080 $Data_Image
