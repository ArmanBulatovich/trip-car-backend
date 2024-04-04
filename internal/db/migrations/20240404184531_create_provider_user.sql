-- +goose Up
-- +goose StatementBegin
CREATE TABLE provider_users (
    id SERIAL PRIMARY KEY,
    provider_id INTEGER REFERENCES providers(id),
    email VARCHAR(32),
    password VARCHAR(256),
    role VARCHAR(32),
    fullname VARCHAR(64),
    phone_number VARCHAR(16),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    created_by INTEGER REFERENCES admins(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE provider_users;
-- +goose StatementEnd
