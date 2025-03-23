# REST API для управления пользователями

REST API для создания, чтения, обновления и удаления пользователей, разработанный на Go.


- Чистая архитектура с разделением ответственности (слои сущностей, репозиториев, сервисов и обработчиков)
- RESTful дизайн API
- База данных PostgreSQL с миграциями
- Контейнеризация с помощью Docker и Docker Compose
- Полное покрытие тестами
- Корректное завершение работы
- Правильная обработка ошибок
- Документация Swagger

## Эндпоинты API

- `GET /api/v1/users` - Получить список пользователей
- `POST /api/v1/users` - Создать нового пользователя
- `GET /api/v1/users/:id` - Получить информацию о пользователе по ID
- `PUT /api/v1/users/:id` - Обновить пользователя по ID
- `DELETE /api/v1/users/:id` - Удалить пользователя по ID

## Документация API

Документация API доступна через Swagger UI по адресу `/swagger/index.html`, когда приложение запущено.


## Запуск приложения

```bash
cp .env.example .env
# Или copy .env.example .env

# Собрать и запустить с помощью Docker Compose
make docker-build
make docker-run

# Остановить контейнеры Docker
make docker-stop
```


## Разработка

```bash
# Запуск тестов
make test

# Генерация Swagger
make swagger
```

---

### Коллекции Postman

В проекте есть `Gravitum.postman_collection.json`, при желании импортировать в Postman.

- **API**
  - **Список пользователей**
    - Метод: GET
    - URL: `http://localhost:8080/api/v1/users`
    - Тело запроса: JSON с данными пользователя
  - **Создать пользователя**
    - Метод: POST
    - URL: `http://localhost:8080/api/v1/users`
    - Тело запроса: JSON с данными пользователя
  - **Получить пользователя**
    - Метод: GET
    - URL: `http://localhost:8080/api/v1/users/:id`
  - **Обновить пользователя**
    - Метод: PUT
    - URL: `http://localhost:8080/api/v1/users/:id`
    - Тело запроса: JSON с обновленными данными
  - **Удалить пользователя**
    - Метод: DELETE
    - URL: `http://localhost:8080/api/v1/users/:id`