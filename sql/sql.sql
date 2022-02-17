CREATE DATABSE IF NOT EXISTS diegobook;

USE diegobook;

DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;

CREATE TABLE users(
  id int auto_increment primary key,
  name varchar(255) not null,
  nickname varchar(255) not null unique,
  email varchar(255) not null unique,
  password varchar(255) not null unique,
  created_at timestamp default current_timestamp
) ENGINE=INNODB;

CREATE TABLE followers(
  user_id int not null,
  FOREIGN KEY (user_id) 
  REFERENCES users(id)
  ON DELETE CASCADE,

  follower_id int not null,
  FOREIGN KEY (follower_id)
  REFERENCES users(id)
  ON DELETE CASCADE,

  primary key(user_id, follower_id)
) ENGINE=INNODB;

CREATE TABLE posts(
  id int auto_increment primary key,
  title varchar(255) not null,
  content varchar(300) not null,

  author_id int not null,
  FOREIGN KEY(author_id)
  REFERENCES users(id)
  ON DELETE CASCADE,

  likes int default 0,
  created_at timestamp default current_timestamp
) ENGINE=INNODB;

INSERT INTO posts(title, content, author_id)
VALUES
("Post user 1", "This is the first post of user 1", 1),
("Post user 2", "This is the first post of user 2", 2),
("Post user 3", "This is the first post of user 3", 3);