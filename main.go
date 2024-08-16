package main

import (
	"belajar-golang/config"
	crud_roles "belajar-golang/crud/roles"
    crud_users  "belajar-golang/crud/users"
	"belajar-golang/entity"
)

func main() {
	config.ConnectDB()

    role := entity.Roles{
        IdRole: 1,
        Role: "Admin",
    }
    
    crud_roles.InsertRoles(config.DB, &role)
    role = entity.Roles{
        IdRole: 2,
        Role: "Tecnician",
    }
    
    crud_roles.InsertRoles(config.DB, &role)
    role = entity.Roles{
        IdRole: 3,
        Role: "Client",
    }
    
    crud_roles.InsertRoles(config.DB, &role)

    user := entity.Users{
        Id: 001,
        Nama: "Fulan",
        Email: "fulan@gmail.com",
        Password: "123456",
        Gender: "L",
        Foto: "Img.jpg",
        Address: "Jl.Cisitu Indah VI No.6",
        RolesIdRole: 1,
    }

    crud_users.InsertUsers(config.DB, &user)
    user = entity.Users{
        Id: 002,
        Nama: "Fulanah",
        Email: "fulanah@gmail.com",
        Password: "123456",
        Gender: "P",
        Foto: "Img.jpg",
        Address: "Jl.Cisitu Indah VI No.6",
        RolesIdRole: 2,
    }

    crud_users.InsertUsers(config.DB, &user)

}