-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    url TEXT NOT NULL,
    UNIQUE(url),
    description TEXT,
    published_at TIMESTAMP NOT NULL,
    feed_id UUID NOT NULL,
    FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);
CREATE INDEX idx_posts_feed_published ON posts (feed_id, published_at DESC);

-- +goose Down
DROP INDEX IF EXISTS idx_posts_feed_published;
DROP TABLE IF EXISTS posts;