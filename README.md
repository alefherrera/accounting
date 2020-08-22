# Accounting notebook

## To run

```
make run
```

## To build

```
make build
```

## Curls

#### Get Balance

```
curl --request GET \
  --url http://localhost:8080/
```

### Get Transactions

```
curl --request GET \
  --url http://localhost:8080/transactions
```

### Get Transaction by Id

```
curl --request GET \
  --url http://localhost:8080/transactions/{id}
```

### Commit transaction

type can be: "credit" or "debit"

```
curl --request POST \
  --url http://localhost:8080/transactions \
  --header 'content-type: application/json' \
  --data '{
	"transaction_type": "credit",
	"amount": 100
}'
```
