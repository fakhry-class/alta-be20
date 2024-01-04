package data

import (
	"fakhry/clean-arch/features/user"

	"gorm.io/gorm"
)

type userQueryRaw struct {
	db *gorm.DB
}

// Update implements user.UserDataInterface.
func (*userQueryRaw) Update(id int, input user.Core) error {
	panic("unimplemented")
}

func NewRaw(db *gorm.DB) user.UserDataInterface {
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
