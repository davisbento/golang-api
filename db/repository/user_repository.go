package repository

import (
	"database/sql"
	"davisbento/golang-api/db/entity"
)

type UserRepository interface {
	FindAll() ([]*entity.User, error)
	FindById(id int) (*entity.User, error)
	Create(user *entity.User) (*entity.User, error)
}

type UserRepositoryImpl struct {
	conn *sql.DB
}

func NewUserRepository(conn *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{conn: conn}
}

func (repo *UserRepositoryImpl) FindAll() ([]*entity.User, error) {
	rows, err := repo.conn.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	var users []*entity.User

	for rows.Next() {
		user := entity.NewUser()

		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo *UserRepositoryImpl) FindById(id int) (*entity.User, error) {
	row := repo.conn.QueryRow("SELECT * FROM users WHERE id = ?", id)

	user := entity.NewUser()

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepositoryImpl) Create(user *entity.User) (*entity.User, error) {
	stmt, err := repo.conn.Prepare("INSERT INTO users (name, email, password) VALUES (?, ?, ?)")

	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(user.Name, user.Email, user.Password)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	user.ID = int(id)

	return user, nil
}
