package repository

import (
	"context"
	"database/sql"
	"penitipan-barang/application/model/domain"
)

type UserRepository interface {
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error)
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, user domain.User)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
}
