package data

import (
	"database/sql"
	"fakhry/clean-arch/features/user"
)

type userQueryRaw struct {
	db *sql.DB
}

// Login implements user.UserDataInterface.
func (*userQueryRaw) Login(email string, password string) (data *user.Core, err error) {
	panic("unimplemented")
}

// Update implements user.UserDataInterface.
func (*userQueryRaw) Update(id int, input user.Core) error {
	panic("unimplemented")
}

func NewRaw(db *sql.DB) user.UserDataInterface {
	return &userQueryRaw{
		db: db,
	}
}

// Insert implements user.UserDataInterface.
func (*userQueryRaw) Insert(input user.Core) error {
	panic("unimplemented")
	//tulis raw sql kalian
}

// SelectAll implements user.UserDataInterface.
func (*userQueryRaw) SelectAll() ([]user.Core, error) {
	panic("unimplemented")
}
