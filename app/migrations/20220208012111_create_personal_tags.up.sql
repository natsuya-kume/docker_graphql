CREATE TABLE personal_tags (
  id         INT NOT NULL AUTO_INCREMENT,
  name       varchar(255) NOT NULL UNIQUE,
  PRIMARY KEY (id)
);