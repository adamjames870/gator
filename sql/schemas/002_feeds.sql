-- +goose Up

CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL, 
    updated_at TIMESTAMP NOT NULL,
    feed_name varchar NOT NULL,
    feed_url varchar UNIQUE NOT NULL,
    created_by_user UUID NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (created_by_user) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down

DROP TABLE feeds;