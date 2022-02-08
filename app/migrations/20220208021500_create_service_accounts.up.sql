CREATE TABLE service_accounts (
  id         INT NOT NULL AUTO_INCREMENT,
  name       varchar(255) NOT NULL UNIQUE,
  email      varchar(255) NOT NULL UNIQUE,
  password   varchar(255) NOT NULL,
  service_id INT NOT NULL,
  PRIMARY KEY (id)
);