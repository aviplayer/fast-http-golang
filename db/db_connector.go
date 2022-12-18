package db

import (
	"context"
	"fast-http-golang/dto"
	"fmt"
	. "github.com/jackc/pgx/v5/pgxpool"
)

var (
	dbString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"db", 5432, "postgres", "postgres", "postgres")
)

type DbConnection struct {
	connection *Pool
}

func CreateConnection() DbConnection {
	dbPool, err := New(context.Background(), dbString)
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}
	return DbConnection{dbPool}
}

func (conn DbConnection) getConnection() (*Conn, error) {
	con, err := conn.connection.Acquire(context.Background())
	return con, err
}

func (conn DbConnection) GetById(id int) (*dto.User, error) {
	ctx := context.Background()
	con, err := conn.getConnection()
	usr := dto.User{}
	if err != nil {
		return nil, err
	}
	defer con.Release()
	err = con.QueryRow(ctx, GET_USER, id).Scan(
		&usr.Id,
		&usr.Username,
		&usr.Email,
		&usr.Age,
	)
	if err != nil {
		return nil, err
	}
	return &usr, nil
}

func (conn DbConnection) InsertUser(name string, email string, age int) (int, error) {
	ctx := context.Background()
	con, err := conn.getConnection()
	if err != nil {
		return -1, err
	}
	defer con.Release()
	var id int
	err = con.QueryRow(ctx, INSERT_USER, name, email, age).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (conn DbConnection) DeleteUser(usr *dto.User) (*dto.User, error) {
	ctx := context.Background()
	con, err := conn.getConnection()
	if err != nil {
		return nil, err
	}
	defer con.Release()
	_, err = con.Query(ctx, DELETE_USER, usr.Id)
	if err != nil {
		return nil, err
	}
	return usr, nil
}
