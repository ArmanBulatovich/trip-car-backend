-- +goose Up
-- +goose StatementBegin
CREATE TABLE admins(
    id SERIAL PRIMARY KEY,
    email VARCHAR(32),
    password VARCHAR(256),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE admins;
-- +goose StatementEnd
