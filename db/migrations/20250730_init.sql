-- migrate:up
CREATE TABLE users (
    user_id UUID NOT NULL,
    first_name VARCHAR(128) NOT NULL,
    last_name VARCHAR(128) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'Asia/Bangkok'),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'Asia/Bangkok'),
    CONSTRAINT PK_users_user_id PRIMARY KEY (user_id)
);

-- migrate:down

