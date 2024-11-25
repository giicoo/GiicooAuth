run:
	
	go build ./cmd/app/main.go
	./main

swagger:
	export PATH=$(go env GOPATH)/bin:$PATH
	swag init -g cmd/app/main.go 
	make run       

docker:
	docker run -it --rm -v ~/Projects/GiicooAuth/:/usr/src/app auth:1.0 