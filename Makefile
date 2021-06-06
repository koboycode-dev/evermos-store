GO := @go
GIN := @gin

goinstall:
	${GO} get .

goprod:
	${GO} build -o main .

gowatch:
	nodemon --exec go run main.go --signal SIGTERM

goclean:
	${GO} clean

gotest:
	${GO} test main_test.go

goformat:
	${GO} fmt ./...