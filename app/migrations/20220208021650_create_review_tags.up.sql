CREATE TABLE review_tags (
  id         INT NOT NULL AUTO_INCREMENT,
  name       varchar(255) NOT NULL UNIQUE,
  user_id INT NOT NULL,
  PRIMARY KEY (id)
);