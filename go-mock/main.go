package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID   int
	Name string
}

type UserRepository interface {
	GetUser(id int) (*User, error)
}

type SQLUserRepository struct {
	db *sql.DB
}

func (r *SQLUserRepository) GetUser(id int) (*User, error) {
	var user User
	err := r.db.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

type UserService struct {
	repo UserRepository
}

func (s *UserService) GetUser(id int) (*User, error) {
	return s.repo.GetUser(id)
}

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO users (id, name) VALUES (1, 'John Doe')")
	if err != nil {
		panic(err)
	}

	userRepo := &SQLUserRepository{db: db}
	userService := &UserService{repo: userRepo}

	user, err := userService.GetUser(1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("User: %+v\n", user)
}
