-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS userdata (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    email VARCHAR(200) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    updated_at TIMESTAMP NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc')
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE userdata;
-- +goose StatementEnd
