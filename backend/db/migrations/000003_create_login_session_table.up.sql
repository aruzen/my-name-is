CREATE TABLE login_sessions
(
    id         UUID PRIMARY KEY,
    user_id    UUID      NOT NULL REFERENCES users (id),
    token      TEXT      NOT NULL, /* hashed */
    expires_at TIMESTAMP NOT NULL DEFAULT (now() + interval '30 min'),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);
