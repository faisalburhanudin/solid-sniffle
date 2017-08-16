CREATE TABLE users
(
  id       INT PRIMARY KEY AUTO_INCREMENT,
  username VARCHAR(50)  NOT NULL,
  password VARCHAR(200) NOT NULL,
  email    VARCHAR(100) NOT NULL
);

CREATE TABLE posts
(
  id      INT AUTO_INCREMENT PRIMARY KEY,
  text    VARCHAR(500) NULL,
  user_id INT          NULL
);

