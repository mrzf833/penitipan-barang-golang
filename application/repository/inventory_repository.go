package repository

import (
	"context"
	"database/sql"
	"penitipan-barang/application/model/domain"
)

type InventoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, inventory domain.Inventory) domain.Inventory
	Update(ctx context.Context, tx *sql.Tx, inventory domain.Inventory) domain.Inventory
	Delete(ctx context.Context, tx *sql.Tx, inventory domain.Inventory)
	FindById(ctx context.Context, tx *sql.Tx, inventoryId int) (domain.Inventory, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Inventory
}
