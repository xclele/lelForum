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
