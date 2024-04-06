package main

import (
	"bwastartup/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//koneksi ke database
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//jika ada error maka cek error nya apa
	if err != nil{
		log.Fatal((err.Error()))
	}
	
	userRepository := user.NewRepository(db)
	
	user := user.User{
		Name: "Test",
	}

	userRepository.Save(user)
}

