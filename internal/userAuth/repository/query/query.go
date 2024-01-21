package query

const CreateUser = `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id;`

const GetUser = `SELECT id,name,email,password from users where email=$1;`

const DeleteUser = `DELETE from users where id=$1;`