CREATE TABLE "public".todos (
    title varchar(128) NOT NULL,
    description varchar(128),
    id numeric NOT NULL,
    CONSTRAINT pk_todos PRIMARY KEY (id)
)


-- create a sql script for creating table if not exist, the columns are id,title, description, user_id?


-- create userTable if not exist with columns id,name,email,password?

CREATE TABLE IF NOT EXISTS userTable (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  email VARCHAR(50) NOT NULL UNIQUE,
  password VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS Posts (
    id INT AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status BOOLEAN,
    user_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES Users(id)
);

