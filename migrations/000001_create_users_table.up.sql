CREATE TABLE IF NOT EXISTS users(
id          SERIAL PRIMARY KEY NOT NULL,
name        varchar(50) NOT NULL,
nickname    varchar(50) NOT NULL,
phone       int NOT NULL,
created_at   TIMESTAMP NOT NULL DEFAULT NOW(),
updated_at   TIMESTAMP
);