package scripts

const UserTable =`CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	email VARCHAR(50) NOT NULL UNIQUE,
	password VARCHAR(50) NOT NULL
  );`


const TodosTable = `CREATE TABLE IF NOT EXISTS todo (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status BOOLEAN,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users(id)
);`