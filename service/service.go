package service

import "database/sql"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserService interface {
	GetUser() (User, error)
	CreateUser(User) (User, error)
	ListUsers() ([]User, error)
}

type defaultUserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) UserService {
	return &defaultUserService{db: db}
}

func (s *defaultUserService) GetUser() (User, error) {
	var u User
	err := s.db.QueryRow(
		"SELECT id, name, email FROM users ORDER BY id LIMIT 1",
	).Scan(&u.ID, &u.Name, &u.Email)
	return u, err
}

func (s *defaultUserService) CreateUser(u User) (User, error) {
	err := s.db.QueryRow(
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id",
		u.Name, u.Email,
	).Scan(&u.ID)
	return u, err
}

func (s *defaultUserService) ListUsers() ([]User, error) {
	rows, err := s.db.Query("SELECT id, name, email FROM users ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, rows.Err()
}
