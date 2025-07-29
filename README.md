### How to run test-backend

Start service container
```sh
docker compose up --build -d
```

Logs service container
```sh
docker compose logs -f
```

Down service container
```sh
docker compose down
```

Vendor: Reinstall vendor packages
```sh
docker compose exec test-backend sh -c "go mod tidy && go mod vendor"	
```

Create network
```sh
docker network create test_network | > /dev/null
```

### How to run migrations

```sh
docker compose up -d
```

### Migration

Create a new migration

```sh
docker compose run --rm customer-migrations new {migration_name} 
```

Run the migrations

```sh
docker compose run --rm customer-migrations
```

### Migration Test

Down database customer-db

```sh
docker compose down customer-db
```

Create a new migration 
```sh
docker compose up -d test-customer-db
```

Run the migrations db test
```sh
docker compose run --rm test-customer-migrations
```

Clean test cache
```sh
go clean -testcache
```
