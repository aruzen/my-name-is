CREATE TABLE "users"
(
    "id"              UUID PRIMARY KEY,
    "username"        VARCHAR(32), /*UNIQUE*/
    "email"           VARCHAR(255) UNIQUE,
    "hashed_password" VARCHAR(255),
    "role"            VARCHAR(20) NOT NULL DEFAULT 'user' CHECK (role IN ('user', 'admin')),
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
)