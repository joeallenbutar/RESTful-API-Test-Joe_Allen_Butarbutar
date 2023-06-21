package configs
 
import (
	"RESTful-API-Test-Joe_Allen_Butarbutar/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB
var err error

func Setup() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading  .env file")
	}

	username 	:= os.Getenv("USERNAME")
	password 	:= os.Getenv("PASSWORD")
	host 		:= os.Getenv("HOST")
	database 	:= os.Getenv("DATABASE")
	port 		:= os.Getenv("PORT")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		username,
		database,
		password,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate([]models.User{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")

}
func GetDB() *gorm.DB {
	return DB
}