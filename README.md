# Accounting notebook

## To run

inside the api folder execute:

```
make run
```

## To build

inside the api folder execute:

```
make build
```

# Web client

## To run

inside the client folder execute:

```
npm start
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
