package crud

import (
	"belajar-golang/entity"

	"github.com/jinzhu/gorm"
)

func InsertUsers(db *gorm.DB, users *entity.Users){
	result := db.Create(users)

	if result.Error != nil{
		panic("Gagal Menambah User")
	}
}