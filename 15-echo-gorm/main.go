package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// struct user gorm model
type User struct {
	gorm.Model
	// ID          uint `gorm:"primaryKey"`
	// CreatedAt   time.Time
	// UpdatedAt   time.Time
	// DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string `json:"name" form:"name"`
	Email       string `gorm:"unique" json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Address     string `json:"address" form:"address"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Role        string `json:"role" form:"role"`
}

/*
TODO 1
buat struct products
id uint
created_at
updated_at
deleted_at
name string
user_id uint FK
description string
*/

var DB *gorm.DB

// database connection
func InitDB() {
	// declare struct config & variable connectionString
	connectionString := os.Getenv("CONNECTION_DB") + "?charset=utf8&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

// db migration
func InitialMigration() {
	DB.AutoMigrate(&User{})
	/*
		TODO 2
		migrate struct product
	*/
}

func main() {
	fmt.Println("running")
	InitDB()
	InitialMigration()

	// create a new echo instance
	e := echo.New()
	// define routes/ endpoint
	e.POST("/users", CreateUserController)
	e.GET("/users", GetAllUserController)
	e.PUT("/users/:user_id", UpdateUserByIdController)

	/*
		TODO 3
		tambahkan endpoint untuk:
		DELETE /users/:user_id
		POST /products
		GET /products
		GET /products/:product_id
		PUT /products/:product_id
		DELETE /products/:product_id
	*/

	//start server and port
	e.Logger.Fatal(e.Start(":8080"))
}

// insert data user
func CreateUserController(c echo.Context) error {
	newUser := User{}
	errBind := c.Bind(&newUser) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data. data not valid",
		})
	}

	// simpan ke DB
	tx := DB.Create(&newUser) // proses query insert
	if tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error insert data. insert failed",
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"message": "insert success",
	})
}

// read data user
func GetAllUserController(c echo.Context) error {
	var usersData []User
	tx := DB.Find(&usersData) // select * from users;
	if tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error read data",
		})
	}
	fmt.Println("users:", usersData)
	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
		"data":    usersData,
	})
}

func UpdateUserByIdController(c echo.Context) error {
	id := c.Param("user_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error. id should be number",
		})
	}
	var userData = User{}
	errBind := c.Bind(&userData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data. data not valid",
		})
	}

	tx := DB.Model(&User{}).Where("id = ?", idParam).Updates(userData)
	if tx.Error != nil {
		// fmt.Println("err:", tx.Error)
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error update " + tx.Error.Error(),
		})
	}

	if tx.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error record not found ",
		})
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
	})
}
