-- +goose Up
-- +goose StatementBegin
CREATE TABLE organization_users (
    id SERIAL PRIMARY KEY,
    organization_id INTEGER REFERENCES organizations(id),
    email VARCHAR(32),
    password VARCHAR(256),
    role VARCHAR(32),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    created_by INTEGER REFERENCES admins(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE organization_users;
-- +goose StatementEnd
