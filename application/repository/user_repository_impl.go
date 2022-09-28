package repository

import (
	"context"
	"database/sql"
	"errors"
	"penitipan-barang/application/helper"
	"penitipan-barang/application/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error) {
	SQL := "SELECT id, name, username, password, role FROM users WHERE username=? LIMIT 1"
	rows, err := tx.QueryContext(ctx, SQL, username)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}

	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.Role)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO users(name, username, password, role) VALUES (?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, user.Name, user.Username, user.Password, user.Role)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE users SET name=?, username=?, password=?, role=? WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Username, user.Password, user.Role, user.Id)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "DELETE FROM users WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "SELECT id, name, username, password, role FROM users WHERE id=? LIMIT 1"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}

	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.Role)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "SELECT id, name, username, password, role FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.Role)
		helper.PanicIfError(err)

		users = append(users, user)
	}

	return users
}
