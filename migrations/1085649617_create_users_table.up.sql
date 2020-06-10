CREATE TABLE users
(
  `id`          int NOT NULL auto increment,
  `name`        varchar(50) NOT NULL,
  `nickname`    varchar(50) NOT NULL,
  `phone`       int(9) NOT NULL,
  `birthday`    DATE NOT NULL,
  `create_at`   DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
  `update_at`   DATETIME ON UPDATE CURRENT_TIMESTAMP NOT NULL,
  primary key (`id`)
) engine = InnoDB
    DEFAULT charset = UTF8;