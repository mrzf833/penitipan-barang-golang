package domain

import (
	"database/sql"
	"time"
)

type Inventory struct {
	Id               int
	CategoryId       sql.NullInt64
	CategoryName     sql.NullString
	DepositName      sql.NullString
	DepositStudentId sql.NullInt64
	DepositAdmin     int
	DepositTime      time.Time
	ItemName         string
	Description      sql.NullString
	Status           string
	TakeName         sql.NullString
	TakeStudentId    sql.NullInt64
	TakeTime         sql.NullTime
	TakeAdmin        sql.NullInt64
}
