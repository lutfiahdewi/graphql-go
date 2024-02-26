# graphql-go

A graphql webservice made using go

# Port

If 8080 blocked, try any port that not in this list
netsh interface ipv4 show excludedportrange protocol=tcp

# Generate model

for newly project:
go run github.com/99designs/gqlgen init

## New Object/Schema

If new object/schema added, then new model need to generated. Run this code:
go run github.com/99designs/gqlgen generate
More: github.com/99designs/gqlgen

# Migrations

## Create table migrations

cd internal/pkg/db/migrations/
migrate create -ext sql -dir mssql -seq create\_[$name]\_table

## No migration in db

use force => edit m.force(n) in db file. n: latest ver +1
OR delete schema_migrations table in th db
check docs in: https://github.com/golang-migrate/migrate

## Run migrations

$ migrate -source file://path/to/migrations -database postgres://localhost:5432/database up 2
example for mysql:
$ migrate -database mysql://root:dbpass@/hackernews -path internal/pkg/db/migrations/mysql up

migrate -database sqlserver://ASUS:12345678@/sobat -path internal/pkg/db/migrations/mssql up
migrate -database sqlserver://ASUS:12345678@localhost:1433/sobat -path internal/pkg/db/migrations/mssql up
sqlserver://ASUS:12345678@localhost?database=sobat

Error migration: Dirty database version 1. Fix and force version. How to fix:
migrate -path PATH_TO_YOUR_MIGRATIONS -database YOUR_DATABASE_URL force VERSION
\*change version with a number of force
migrate -database sqlserver://ASUS:12345678@localhost:1433/sobat -path internal/pkg/db/migrations/mssql up force 5

# Auth

##

# Query/Mutations example

## Dateformat

Full length :
"2006-01-02T15:04:05.999999999Z" (UTC )
"2006-01-02T15:04:05.999999999+03:30"	(A DateTime with +3h 30min offset)

## Query

## Mutations

### Create

mutation {
createUser(input: {username: "user123", password: "123"})
}
mutation {
createKegiatan(input: {title: "kegiatan 1", status: "jalan"}) {
title
status
}
}

### Update

### Delete

# Prevented not logged in users from submitting links => Set the Auth Header. Example:

{
"Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
}

# Note for programmer
## sql statement
## model for the return argument
Model cannot accept nil
