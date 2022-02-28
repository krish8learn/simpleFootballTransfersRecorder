sqlcgen:
	sqlc generate 

sqlccompile:
	sqlc compile

test:
	go test -v cover ./...

gomod:
	go mod tidy -compat=1.17	


.PHONY:  sqlcgen sqlccompile test gomod
