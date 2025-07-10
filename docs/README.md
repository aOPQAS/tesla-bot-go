# Tesla Bot Go

Tesla Bot Go — это Telegram-бот и REST API для удалённого управления автомобилем Tesla. Проект разработан на языке Go с использованием фреймворка Fiber и PostgreSQL в качестве базы данных.

## Стек технологий

- Go — основной язык программирования
- Fiber — веб-фреймворк для REST API
- PostgreSQL — СУБД для хранения данных
- Telegram Bot API — для взаимодействия с пользователем
- JWT (Access + Refresh) — аутентификация
- Docker — контейнеризация и деплой

## Возможности

- Аутентификация через Telegram
- Генерация и обновление JWT-токенов
- Управление функциями Tesla через REST и Telegram
- Разделение на API и Telegram-бот
- Безопасное хранение refresh-токенов в базе данных

## Запуск проекта

1. Клонируйте репозиторий
   git clone https://github.com/yourusername/tesla-bot-go.git
   cd tesla-bot-go
