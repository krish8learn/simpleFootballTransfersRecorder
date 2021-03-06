DB_URL=postgresql://root:secret@localhost:5432/simple_transfers?sslmode=disable

postgres start:
	sudo docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	sudo docker exec -it postgres createdb --username=root --owner=root simple_transfers

dropdb:
	sudo docker exec -it postgres dropdb --username=root --owner=root simple_transfers

migrateup:
	migrate -path DB/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path DB/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path DB/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path DB/migration -database "$(DB_URL)" -verbose down 1

sqlcgen:
	sqlc generate 

sqlccompile:
	sqlc compile

test:
	go test -v -cover ./...

mock:
	mockgen -package mockdb -destination DB/mock/tranaction.go github.com/krish8learn/simpleFootballTransfersRecorder/DB/sqlc Transaction

gomod:
	go mod tidy -compat=1.17	

.PHONY: postgres createdb dropdb migrateup migratedown sqlcgen sqlccompile test gomod mock