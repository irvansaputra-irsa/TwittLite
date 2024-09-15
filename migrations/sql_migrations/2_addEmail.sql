-- +migrate Up
-- +migrate StatementBegin
ALTER TABLE users ADD email VARCHAR(255);

-- +migrate StatementEnd