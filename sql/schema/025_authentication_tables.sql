-- Needed for authjs authentication

-- +goose Up
CREATE TABLE verification_token
(
    identifier TEXT NOT NULL,
    expires TIMESTAMPTZ NOT NULL,
    token TEXT NOT NULL,

    PRIMARY KEY (identifier, token)
);

CREATE TABLE accounts
(
    id SERIAL,
    "userId" INTEGER NOT NULL,
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

CREATE TABLE session_state(
    id SERIAL,
    "userId" INTEGER NOT NULL,
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
DROP TABLE session_state;

ALTER TABLE users
DROP COLUMN "emailVerified",
DROP COLUMN name,
DROP image;