CREATE TABLE users (
  id         INT NOT NULL AUTO_INCREMENT,
  name       varchar(255) NOT NULL UNIQUE,
  service_id INT NOT NULL,
  PRIMARY KEY (id)
);