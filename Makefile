db-init:
	psql -c 'CREATE DATABASE "product-management"' -U $(user)
migrationup:
	migrate -path db/migration -database "postgres://$(user):$(password)@$(host):$(port)/product-management?sslmode=disable" -verbose up
migrationdown:
	migrate -path db/migration -database "postgres://$(user):$(password)@$(host):$(port)/product-management?sslmode=disable" -verbose down
