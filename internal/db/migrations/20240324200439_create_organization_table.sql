-- +goose Up
-- +goose StatementBegin
CREATE TABLE organizations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64),
    slug VARCHAR(32) UNIQUE,
    metadata JSON,
    created_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP,
    created_by INTEGER REFERENCES admins(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE organizations;
-- +goose StatementEnd
