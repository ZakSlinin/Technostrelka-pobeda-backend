# Subscriptions Service API

## Сервис управления подписками пользователей. Все запросы проходят через API Gateway и используют формат JSON.

### Общие сведения

Base URL: /api/subscriptions

Формат ID: Все user_id и subscription_id должны быть в формате UUID v4.

Стандартный ответ при ошибке:
    
```JSON
{
"error": "INTERNAL_ERROR",
"message": "Описание ошибки для разработчика",
"timestamp": "2026-03-01T12:00:00Z"
}
```

### Endpoints

1. Создание подписки
   Создает новую подписку для конкретного пользователя.

Метод: POST

/api/subscriptions/create/:user_id


```JSON
{
"name": "Netflix",
"cost": 12.99,
"next_billing": "2026-04-01T00:00:00Z",
"category": "Entertainment"
}
```

Успех: 201 Created + Объект созданной подписки.

2. Обновление подписки
   Частичное или полное обновление данных существующей подписки.

Метод: POST (или PATCH)

/api/subscriptions/update/:user_id/:subscription_id

Body: Принимает поля из UpdateSubscriptionRequest.

Успех: 200 OK

```JSON
{ "message": "subscription updated successfully" }
Ошибки: 404 Not Found — если подписка с таким ID не найдена у данного пользователя.
```

400 Bad Request — если cost < 0.

3. Получение всех подписок пользователя
   Возвращает массив всех подписок, привязанных к user_id.

Метод: GET

Путь: /api/subscriptions/all/:user_id

Body: Не требуется.

Успех: 200 OK + [...] (массив объектов подписок).

4. Удаление подписки
   Безвозвратное удаление записи из базы данных.

Метод: DELETE (или POST)

/api/subscriptions/delete/:user_id/:subscription_id

Успех: 200 OK

```JSON
{ "message": "subscription deleted successfully" }
```