-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    first_name VARCHAR(25) NOT NULL,
	last_name VARCHAR(25) NOT NULL,
	surname VARCHAR(25) NOT NULL,
	avatar_url VARCHAR(255),
	status VARCHAR(15),
	date_of_birth DATE NOT NULL,
	created_at TIMESTAMPTZ NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS users;
-- +goose StatementEnd
