GO := @go
GIN := @gin

goinstall:
	${GO} get .

goprod:
	${GO} build -o main .

gowatch:
	nodemon --exec go run main.go --signal SIGTERM

goclean:
	go clean

gotest:
	${GO} test -v

goformat:
	${GO} fmt ./...