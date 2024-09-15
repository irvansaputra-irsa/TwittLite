-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE
    users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(256) NOT NULL,
        password VARCHAR(256) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        created_by VARCHAR(356) DEFAULT 'SYSTEM',
        modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        modified_by VARCHAR(356) DEFAULT 'SYSTEM'
    );

CREATE TABLE
    follows (
        id SERIAL PRIMARY KEY,
        follower_id INTEGER NOT NULL REFERENCES users (id),
        following_id INTEGER NOT NULL REFERENCES users (id),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        created_by VARCHAR(356) DEFAULT 'SYSTEM',
        modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        modified_by VARCHAR(356) DEFAULT 'SYSTEM'
    );

CREATE TABLE
    posts (
        id SERIAL PRIMARY KEY,
        content VARCHAR(256) NOT NULL,
        user_id INTEGER NOT NULL REFERENCES users (id),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        created_by VARCHAR(356) DEFAULT 'SYSTEM',
        modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        modified_by VARCHAR(356) DEFAULT 'SYSTEM'
    );

CREATE TABLE
    comments (
        id SERIAL PRIMARY KEY,
        content VARCHAR(256) NOT NULL,
        user_id INTEGER NOT NULL REFERENCES users (id),
        post_id INTEGER NOT NULL REFERENCES posts (id),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        created_by VARCHAR(356) DEFAULT 'SYSTEM',
        modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        modified_by VARCHAR(356) DEFAULT 'SYSTEM'
    );

-- +migrate StatementEnd