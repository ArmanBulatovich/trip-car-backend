# trip-car-backend
## Database
### migrate
    cd {TRIP_ROOT}/internal/db/migrations
    goose postgres "postgresql://user:password@127.0.0.1:5432/tripcar?sslmode=disable" up

## Run
    cd {TRIP_ROOT}/cmd
    go run main.go
