package main

import (
	"bwastartup/user"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

	fmt.Println("Connection to database is good")

	var users []user.User
	db.Find(&users)

	for _, user := range users {
		fmt.Println(user.Name)
		fmt.Println(user.Email)
	}

	router := gin.Default()
	router.GET("/handler", handler)
	router.Run()
}

func handler(c *gin.Context){
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//jika ada error maka cek error nya apa
	if err != nil{
		log.Fatal((err.Error()))
	}

	var users []user.User
	db.Find(&users)

	c.JSON(http.StatusOK, users)

}