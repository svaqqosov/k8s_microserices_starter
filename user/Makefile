GO := @go
NODEMON := @nodemon

goinstall:
	${GO} get .

godev:
	${NODEMON} --watch './**/*.go' --signal SIGTERM --exec 'go' run ./main.go

goprod:
	${GO} build -o main .

gotest:
	${GO} test -v

goformat:
	${GO} fmt ./...

.PHONY: cert
cert:
	openssl genrsa -out cert/id_rsa 4096
	openssl rsa -in cert/id_rsa -pubout -out cert/id_rsa.pub