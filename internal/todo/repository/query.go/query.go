package query

const CreateTodo = `INSERT INTO todo ( title, description, status) VALUES ( $1, $2, $3 ) RETURNING id;`
const GetTodo = `SELECT
id, title, description, status
FROM
"public".todo o;`

const UpdateTodoStatus = `UPDATE todo SET status=$2 WHERE id=$1`
const UpdateTodoData = `UPDATE "public".todo SET title=$1,description=$2 WHERE id=$3;`

// select todo by Id?
const GetTodoById = `SELECT
id, title, description, status
FROM
"public".todo WHERE id=$1;`

// delete Todo
const DeleteTodo = `DELETE FROM todo where id=$1`
