package entity

import "time"

type Roles struct {
	IdRole int    `gorm:"column:id_role;primaryKey"`
	Role   string `gorm:"column:role"`
}

type Users struct {
	Id        int    `gorm:"column:id;primaryKey;autoIcrement"`
	Nama      string `gorm:"column:name"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
	Gender    string `gorm:"column:gender"`
	Foto      string `gorm:"column:foto"`
	Address   string `gorm:"column:address"`
	RolesRole string `gorm:"column:role"`
	Roles     Roles  `gorm:"foreignKey:RolesRole"`
}

type Ac struct {
	IdAc   int     `gorm:"column:id_ac;primaryKey;autoIncrement"`
	NameAc string  `gorm:"column:name"`
	Brand  float32 `gorm:"column:pk"`
	Price  int     `gorm:"column:price"`
}

type Services struct {
	IdService    int 		`gorm:"column:id_service;primaryKey;autoIncrement"`
	IdTechnician int 		`gorm:"column:id_technician"`
	IdClient     int 		`gorm:"column:id_client"`
	AcIdAc       int 		`gorm:"column:id_ac"`
	Ac           Ac  		`gorm:"foreignKey:AcIdAc"`
	Date         time.Time 	`gorm:"date"`
	Status		 string		`gorm:"status"`
}
