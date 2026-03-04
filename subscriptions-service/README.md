### Subscriptions endpoints 

## Все эндпоинты сгруппированы по /api/subscriptions (добавляй в начало)
- **Create**: `POST /create` + Header `X-User-Id`
- **List All**: `GET /all` + Header `X-User-Id`
- **Update**: `POST /update/:subscription_id` + Header `X-User-Id`
- **Delete**: `DELETE /delete/:subscription_id` + Header `X-User-Id`
- **Subscription**  Get by ID** `GET /:subscription_id` + Header `X-User-ID`

*Сервис сам проверяет, принадлежит ли `:subscription_id` пользователю из хедера.*