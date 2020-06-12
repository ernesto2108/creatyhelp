CREATE TABLE IF NOT EXISTS users(
    id          serial primary key,
    name        varchar(50) NOT NULL,
    nickname    varchar(50) NOT NULL,
    phone       int NOT NULL,
    create_at   TIMESTAMP  NOT NULL,
    update_at   TIMESTAMP ,
    PRIMARY KEY (id)
  )
