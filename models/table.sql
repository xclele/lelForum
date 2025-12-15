CREATE TABLE "user" (
    id BIGSERIAL PRIMARY KEY,
    user_id NUMERIC(20,0) NOT NULL,
    username VARCHAR(64) NOT NULL,
    password VARCHAR(64) NOT NULL,
    email VARCHAR(64),
    gender SMALLINT NOT NULL DEFAULT 0,
    create_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT idx_username UNIQUE (username),
    CONSTRAINT idx_user_id UNIQUE (user_id)
);

DROP TABLE IF EXISTS community;
CREATE TABLE community (
    id SERIAL PRIMARY KEY,
    community_id INTEGER NOT NULL,
    community_name VARCHAR(128) NOT NULL,
    introduction VARCHAR(256) NOT NULL,
    create_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT idx_community_id UNIQUE (community_id),
    CONSTRAINT idx_community_name UNIQUE (community_name)
);
INSERT INTO community VALUES (1, 1, 'Go', 'Golang', '2016-11-01 08:10:10', '2016-11-01 08:10:10');

DROP TABLE IF EXISTS post;
CREATE TABLE post (
    id BIGSERIAL PRIMARY KEY,
    post_id BIGINT NOT NULL,
    title VARCHAR(128) NOT NULL,
    content TEXT NOT NULL,
    author_id BIGINT NOT NULL,
    community_id INTEGER NOT NULL,
    status SMALLINT NOT NULL DEFAULT 1,
    create_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_post_id ON post(post_id);
CREATE INDEX IF NOT EXISTS idx_author_id ON post(author_id);
CREATE INDEX IF NOT EXISTS idx_community_id ON post(community_id);
