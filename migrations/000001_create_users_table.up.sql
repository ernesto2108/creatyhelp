CREATE TABLE IF NOT EXISTS users(
id          serial PRIMARY KEY NOT NULL,
name        varchar(50) NOT NULL,
nickname    varchar(50) NOT NULL,
phone       int NOT NULL,
create_at   TIMESTAMP,
update_at   TIMESTAMP
);