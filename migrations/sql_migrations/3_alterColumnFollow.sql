-- +migrate Up
-- +migrate StatementBegin
ALTER TABLE follows
RENAME COLUMN followed_id TO following_id;

-- +migrate StatementEnd