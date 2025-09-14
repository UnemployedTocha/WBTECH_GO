# Демонстрационный сервис

Сервис с простейшим интерфейсом, отображающий данные о заказе

---

## Демонстрация работы

https://drive.google.com/drive/folders/13cS5Rs2hZZD6OFThLxIg_Vs9kUqtYlpJ?usp=sharing

---

## Задание

[Формулировка задания](./task.md)

---

## Требования

- [Go](https://golang.org/) >= 1.20
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

## Сборка и запуск через Docker Compose

1. Склонируйте репозиторий:

```bash
git clone https://github.com/UnemployedTocha/WBTECH_GO.git && cd WBTECH_GO/L0/
```

2. Заполните .env конфиг файл. Пример который можно использовать:
```
POSTGRES_USER=user
POSTGRES_PASSWORD=qwerty
POSTGRES_DB=orders_db
POSTGRES_HOST=db
POSTGRES_PORT=5436
POSTGRES_INTERNAL_PORT=5432
SSL_MODE=disable

SERVICE_PORT=8098
SERVICE_INTERNAL_PORT=8089
```

3. Соберите и запустите контейнеры:
```bash
docker-compose up --build
```