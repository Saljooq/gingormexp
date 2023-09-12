package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CreatedUser struct {
	gorm.Model
	Name  string
	Email string

	// person_id SERIAL,
	// name varchar(80) NOT NULL,
	// nickname varchar(50) NOT NULL,
	// email varchar(80) NOT NULL,
	// password varchar(80) NOT NULL,
}

var DB *gorm.DB

func OpenConnection() {

	host := "postgres"
	port := 5432
	user := "postgres"
	password := "postgres"

	psql_info := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s sslmode=disable",
		host, port, user, password)

	db, err := gorm.Open(postgres.Open(psql_info), &gorm.Config{})
	if err != nil {
		log.Fatal("Couldn't connect to the db :(")
		panic(err)
	}

	DB = db
}

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	OpenConnection()

}

func main() {
	r := gin.Default()

	DB.AutoMigrate(&CreatedUser{})

	r.GET("/post", func(c *gin.Context) {

		usr := CreatedUser{Name: "Sal", Email: "sal@e.com"}
		DB.Create(&usr)

		log.Printf("User created successfully. id: %v", usr.ID)

		c.JSON(200, gin.H{
			"message": "successful",
		})
	})

	r.GET("/all", func(c *gin.Context) {

		var usrs []CreatedUser

		DB.Find(&usrs)

		c.JSON(200, gin.H{"res": usrs})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:3000
}
