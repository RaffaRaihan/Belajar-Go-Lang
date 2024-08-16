package crud

import (
	"belajar-golang/entity"

	"github.com/jinzhu/gorm"
)

func InsertRoles(db *gorm.DB, roles *entity.Roles) {
	result := db.Create(roles)

	if result.Error != nil{
		panic("Gagal Menambah Role")
	}
}