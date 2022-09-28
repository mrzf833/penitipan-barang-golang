package repository

import (
	"context"
	"database/sql"
	"errors"
	"penitipan-barang/application/helper"
	"penitipan-barang/application/model/domain"
)

type StudentRepositoryImpl struct {
}

func NewStudentRepository() StudentRepository {
	return &StudentRepositoryImpl{}
}

func (repository *StudentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, student domain.Student) domain.Student {
	SQL := "INSERT INTO students(fullname,email,phone_number,address) VALUES (?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, student.Fullname, student.Email, student.PhoneNumber, student.Address)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	student.Id = int(id)
	return student
}

func (repository *StudentRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, student domain.Student) domain.Student {
	SQL := "UPDATE students SET fullname=?, email=?, phone_number=?, address=? WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, student.Fullname, student.Email, student.PhoneNumber, student.Address, student.Id)
	helper.PanicIfError(err)

	return student
}

func (repository *StudentRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, student domain.Student) {
	SQL := "DELETE FROM students WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, student.Id)
	helper.PanicIfError(err)
}

func (repository *StudentRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, studentId int) (domain.Student, error) {
	SQL := "SELECT id, fullname, email, phone_number, address FROM students WHERE id=? LIMIT 1"
	rows, err := tx.QueryContext(ctx, SQL, studentId)
	helper.PanicIfError(err)
	defer rows.Close()

	student := domain.Student{}

	if rows.Next() {
		err = rows.Scan(&student.Id, &student.Fullname, &student.Email, &student.PhoneNumber, &student.Address)
		helper.PanicIfError(err)
		return student, nil
	} else {
		return student, errors.New("student is not found")
	}
}

func (repository *StudentRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Student {
	SQL := "SELECT id, fullname, email, phone_number, address FROM students"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var students []domain.Student

	for rows.Next() {
		student := domain.Student{}
		err := rows.Scan(&student.Id, &student.Fullname, &student.Email, &student.PhoneNumber, &student.Address)
		helper.PanicIfError(err)

		students = append(students, student)
	}

	return students
}
