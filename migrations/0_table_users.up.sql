CREATE TABLE IF NOT EXISTS users(
  `id`          int NOT NULL auto_increment,
  `name`        varchar(50) NOT NULL,
  `nickname`    varchar(50) NOT NULL,
  `phone`       int(9) NOT NULL,
  `create_at`   DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
  `update_at`   DATETIME ON UPDATE CURRENT_TIMESTAMP NOT NULL,
  PRIMARY KEY (`id`)
)engine = InnoDB
  DEFAULT charset = utf8;