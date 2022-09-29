# libro

A library management system built with Go.

## Setup

```console
$ createdb libro
$ psql libro < tables.sql
$ go run ./cmd/web --dbName libro --dbHost localhost --dbPort 5432 --dbUser libro_user --dbPassword libro_password 
```