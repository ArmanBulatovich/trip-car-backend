# trip-car-backend
## Database
### migrate
    cd {TRIP_ROOT}/internal/db/migrations
    goose postgres "postgresql://user:password@127.0.0.1:5432/tripcar?sslmode=disable" up
## Create super admin
    go run main.go createsuperadmin
## Run
    go run main.go
