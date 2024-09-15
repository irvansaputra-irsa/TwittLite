-- +migrate Up
-- +migrate StatementBegin
ALTER TABLE users ADD bio varchar(255);

ALTER TABLE users ADD location varchar(255);

-- +migrate StatementEnd