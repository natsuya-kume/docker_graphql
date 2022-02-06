CREATE TABLE admins (
  id         bigint(20) NOT NULL AUTO_INCREMENT,
  name       varchar(255) NOT NULL UNIQUE,
  email      varchar(255) NOT NULL UNIQUE,
  password   varchar(255) NOT NULL,
  PRIMARY KEY (id)
);