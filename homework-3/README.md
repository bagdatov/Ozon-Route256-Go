# homework-3

Домашнее задание к мастер-классу "Проектирование распределенных систем"
Третье задание будет состоять из нескольких частей. И первая, как и в прошлый раз, связана с проектированием. В этот раз я предлагаю вам подумать как вообще выглядит Ozon. 

Вам потребуется нарисовать крупномасштабную схему, аналогичной той, что мы делали на мастер-классе для сервиса заказа такси. Не обязательно прописывать всё, что есть, постарайтесь показать как вы видите такую систему, как Озон. Уровень детализации определяете вы сами (как и мы с сервисом заказа такси).

После этого вам надо будет выбрать минимум три сервиса, которые связаны общей транзакцией. Такой транзакцией может быть оплата заказа, выдача товара, комплектация отправления. Для этих сервисов вам понадобиться расписать детальную реализацию, аналогичную той, что мы делали для сервиса сокращения ссылок.

Ссылка на наши схемы с мастер класса: https://miro.com/app/board/uXjVOzhjoyc=/?share_link_id=201558986951

Для сдачи задания, создайте новый репозиторий и сделайте MR либо с ссылками на схемы, либо с самими схемами, чтобы обсуждать MR можно было в гитлабе. 

Ещё раз:
- Озон в целом
- Минимум три сервиса с расчётом нагрузки, хранилищ, кэшей и всего остального.


-------

## Implementation Concept:

### Ссылка просто посмотреть на доску: https://miro.com/app/board/uXjVOxSlqWs=/?share_link_id=853568345782

### Ссылка на редактирование доски: https://miro.com/welcomeonboard/QXpOVkpteEpwSmxhdmI3Q3NpYXB3WE9jTEZ0R2ZtNEFRbWJ5Qmg2Y3pnUEREMGoxYnpNSFYyOXF1ZERJVGI3V3wzNDU4NzY0NTI2Mzc0MTAwMjcw?share_link_id=173071933586


-----

## Implementation

Хореографическая сага через очереди сообщений.
Чтобы поднять приложение необходим `docker-compose`. 
Зависимости скачиваются долговато.

Чтобы реализовать партицирование пришлось установить статичные ip адреса контейнерам.
Кэширование реализовано через Redis.

Топики в кафке:
  - `incoming` новые заказы
  - `reservation` подтверждения бронирования
  - `reset` откаты при ошибках

Kafka UI: `localhost:8080`

### service-create-order
- Создает заказы и публикует в очередь `incoming`
- `localhost:8000`
- API:
  - `/api/v1/orders/create` создать заказ
  - `/api/v1/orders/cancel` отменить заказ (компенсирующая транзакция)
  - `/metrics` метрики

### service-monitoring-order
- Мониторит статус заказа
- Партицирование статусов заказов (shard1, shard2)
- `localhost:8001`
- API:
    - `/api/v1/status` статус заказа
    - `/metrics` метрики

### service-store
- Поворотная транзакция при неуспешном резервировании
- Кэширование статусов резерва (Redis), время жизни кэша 15 мин по умолчению
- `localhost:8002`
- API:
    - `/api/v1/items/find` найти товар и проверить его резервирование
    - `/api/v1/items/add` создать 
    - `/metrics` метрики

Примеры запросов есть Insomnia_requests.json
Или ниже curl запросы.

### curl запросы:
```bash
echo "1. Create Order"
curl --request POST \
  --url http://localhost:8000/api/v1/orders/create \
  --header 'Content-Type: application/json' \
  --data '{
	"item_id": 1,
	"seller_id": 1,
	"client_id": 1
}'

echo "2. Monitoring Status"
curl --request GET \
  --url 'http://localhost:8001/api/v1/status?order_id=1'

echo "3. Check Item Reservations"
curl --request GET \
  --url 'http://localhost:8002/api/v1/items/find?item_id=1'

echo "4. Cancel Request"
curl --request POST \
  --url 'http://localhost:8000/api/v1/orders/cancel?order_id=1'
```