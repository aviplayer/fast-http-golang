package db

const GET_USER = `SELECT * FROM users as u
WHERE u.id=$1`

const INSERT_USER = `INSERT INTO users (name, email, age)
values ($1, $2, $3) RETURNING id`

const DELETE_USER = `DELETE FROM users as u
WHERE u.id=$1`
