CREATE TABLE hue_records
(
    id         UUID PRIMARY KEY,
    user_name  varchar(255) NOT NULL,
    choices    JSONB        NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);
