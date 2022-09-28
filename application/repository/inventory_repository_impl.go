package repository

import (
	"context"
	"database/sql"
	"errors"
	"penitipan-barang/application/helper"
	"penitipan-barang/application/model/domain"
)

type InventoryRepositoryImpl struct {
}

func NewInventoryRepository() InventoryRepository {
	return &InventoryRepositoryImpl{}
}

func (repository *InventoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, inventory domain.Inventory) domain.Inventory {
	status := "deposit"
	SQL := "INSERT INTO inventories(category_id, deposit_student_id, deposit_name, deposit_time, item_name, description, deposit_admin, status) VALUES (?,?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL,
		inventory.CategoryId, inventory.DepositStudentId, inventory.DepositName,
		inventory.DepositTime, inventory.ItemName, inventory.Description,
		inventory.DepositAdmin,
		status,
	)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	inventory.Id = int(id)
	inventory.Status = status
	return inventory
}

func (repository *InventoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, inventory domain.Inventory) domain.Inventory {
	SQL := `UPDATE inventories SET 
		category_id=?, deposit_student_id=?, deposit_name=?, 
		deposit_time=?, item_name=?, description=?,
		status=?, take_student_id=?, take_name=?,
		take_time=?
		WHERE id=?`
	_, err := tx.ExecContext(ctx, SQL,
		inventory.CategoryId, inventory.DepositStudentId, inventory.DepositName,
		inventory.DepositTime, inventory.ItemName, inventory.Description,
		inventory.Status, inventory.TakeStudentId, inventory.TakeName,
		inventory.TakeTime,
		inventory.Id,
	)
	helper.PanicIfError(err)

	return inventory
}

func (repository *InventoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, inventory domain.Inventory) {
	SQL := "DELETE FROM inventories WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, inventory.Id)
	helper.PanicIfError(err)
}

func (repository *InventoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, inventoryId int) (domain.Inventory, error) {
	SQL := `SELECT 
				invt.id, invt.category_id, cg.name, invt.deposit_name,
				invt.deposit_student_id, invt.deposit_admin, invt.deposit_time, invt.item_name, 
				invt.description, invt.status, invt.take_name, invt.take_student_id,
				invt.take_time, invt.take_admin
			FROM inventories AS invt
			LEFT JOIN categories AS cg
			ON invt.category_id = cg.id
			WHERE invt.id=? LIMIT 1`
	rows, err := tx.QueryContext(ctx, SQL, inventoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	inventory := domain.Inventory{}

	if rows.Next() {
		err = rows.Scan(
			&inventory.Id, &inventory.CategoryId, &inventory.CategoryName, &inventory.DepositName,
			&inventory.DepositStudentId, &inventory.DepositAdmin, &inventory.DepositTime, &inventory.ItemName,
			&inventory.Description, &inventory.Status, &inventory.TakeName, &inventory.TakeStudentId,
			&inventory.TakeTime, &inventory.TakeAdmin)
		helper.PanicIfError(err)
		return inventory, nil
	} else {
		return inventory, errors.New("inventory is not found")
	}
}

func (repository *InventoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Inventory {
	SQL := `SELECT 
				invt.id, invt.category_id, cg.name, invt.deposit_name,
				invt.deposit_student_id, invt.deposit_admin, invt.deposit_time, invt.item_name, 
				invt.description, invt.status, invt.take_name, invt.take_student_id,
				invt.take_time, invt.take_admin
			FROM inventories AS invt
			LEFT JOIN categories AS cg
			ON invt.category_id = cg.id`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var inventories []domain.Inventory

	for rows.Next() {
		inventory := domain.Inventory{}
		err := rows.Scan(
			&inventory.Id, &inventory.CategoryId, &inventory.CategoryName, &inventory.DepositName,
			&inventory.DepositStudentId, &inventory.DepositAdmin, &inventory.DepositTime, &inventory.ItemName,
			&inventory.Description, &inventory.Status, &inventory.TakeName, &inventory.TakeStudentId,
			&inventory.TakeTime, &inventory.TakeAdmin,
		)
		helper.PanicIfError(err)

		inventories = append(inventories, inventory)
	}

	return inventories
}
