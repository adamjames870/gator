-- +goose Up

CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL, 
    updated_at TIMESTAMP NOT NULL,
    feed_name varchar NOT NULL,
    feed_url varchar UNIQUE NOT NULL,
    user_id UUID NOT NULL,
    CONSTRAINT fk_users FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down

DROP TABLE feeds;