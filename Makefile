install:
	go get -u github.com/gorilla/mux
	go get -u github.com/gorilla/handlers
	go get github.com/jmoiron/sqlx

run:
	go run main.go
