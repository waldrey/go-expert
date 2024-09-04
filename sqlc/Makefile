createmigration:
	migrate create -ex=sql -dir=sqlc/migrations -seq init

migrate:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up 

migratedown:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down 

.PHONY: migrate migratedown migrate