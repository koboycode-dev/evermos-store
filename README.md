## Store

1. Explain what you think happened that caused the bad review during our 12.12 event and why it happened.
   
   Because there is a 12.12 event and orders are busy, and many orders simultaneously, when stock orders are not updated or not sequentially, in the sense of not    queuing one by one, who is the first to order and the next, so there is an error in the stock.

   there are several approaches to this solution
   first will use the queue system when customers place orders, can use redis etc.

   secondly check the stock quantity when going to checkout ##

2. 

## How to run

copy environment variable in terminal
```sh
cp .env-sample .env
```

fill the `DATABASE_URI_DEV` and JWT_SECRET

run manualy

```sh
go run main.go
```

testing
```sh
go test main_test.go
```


or run via makefile
for watching development mode
```sh
make gowatch
```

for test
```
make gotest
```
