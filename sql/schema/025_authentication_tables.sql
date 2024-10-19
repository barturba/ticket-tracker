-- Needed for authjs authentication

-- +goose Up
ALTER TABLE users ALTER COLUMN id SET DEFAULT gen_random_uuid();
ALTER TABLE users ALTER COLUMN created_at SET DEFAULT now();
ALTER TABLE users ALTER COLUMN updated_at SET DEFAULT now();

CREATE TABLE verification_token
(
    identifier TEXT NOT NULL,
    expires TIMESTAMPTZ NOT NULL,
    token TEXT NOT NULL,

    PRIMARY KEY (identifier, token)
);

CREATE TABLE accounts
(
    id uuid default gen_random_uuid(),
    "userId" uuid NOT NULL,
    type VARCHAR(255) NOT NULL,
    provider VARCHAR(255) NOT NULL,
    "providerAccountId" VARCHAR(255) NOT NULL,
    refresh_token TEXT,
    access_token TEXT,
    expires_at BIGINT,
    id_token TEXT,
    scope TEXT,
    session_state TEXT,
    token_type TEXT,

    PRIMARY KEY (id)
);

CREATE TABLE sessions
(
  id uuid default gen_random_uuid(),
  "userId" uuid NOT NULL,
  expires TIMESTAMPTZ NOT NULL,
  "sessionToken" VARCHAR(255) NOT NULL,
 
  PRIMARY KEY (id)
);
 

ALTER TABLE users
ADD "emailVerified" TIMESTAMPTZ,
ADD name TEXT,
ADD image TEXT;

-- +goose Down

DROP TABLE verification_token;
DROP TABLE accounts;
DROP TABLE sessions;

ALTER TABLE users
DROP COLUMN "emailVerified",
DROP COLUMN name,
DROP image;

ALTER TABLE users ALTER COLUMN id DROP DEFAULT;
ALTER TABLE users ALTER COLUMN created_at DROP DEFAULT;
ALTER TABLE users ALTER COLUMN updated_at DROP DEFAULT;