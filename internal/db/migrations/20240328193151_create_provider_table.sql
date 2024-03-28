-- +goose Up
-- +goose StatementBegin
CREATE TABLE providers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64),
    slug VARCHAR(32) UNIQUE,
    metadata JSON,
    bin VARCHAR(32),
    email VARCHAR(32),
    phone_number VARCHAR(32),
    created_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP,
    created_by INTEGER REFERENCES admins(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE providers;
-- +goose StatementEnd
