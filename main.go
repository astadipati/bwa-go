package main

import (
	"bwastartup/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "mase:Kul0nuwun@tcp(127.0.0.1:3306)/bwa_startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	//kita panggil repo
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisteruserInput{}
	userInput.Name = "name dari service"
	userInput.Email = "service@gmail.com"
	userInput.Occupation = "Anak band"
	userInput.Password = "password"

	userService.RegisterUser(userInput)
}

// func handler(c *gin.Context) {
// 	dsn := "mase:Kul0nuwun@tcp(127.0.0.1:3306)/bwa_startup?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	var users []user.User
// 	db.Find(&users)

// 	c.JSON(http.StatusOK, users)
// }
