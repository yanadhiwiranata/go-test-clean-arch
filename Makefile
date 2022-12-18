BINARY=engine

run:
	go run cli/main.go

docker-run:
	docker-compose up --build -d

docker-stop:
	docker-compose down

unittest:
	go test ./...

engine:
	go build -o ${BINARY} cli/*.go

docker:
	docker build -t go-test-clean-arch .

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi


.PHONY: clean install unittest build docker run docker-run docker-stop