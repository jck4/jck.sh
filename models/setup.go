package models
import (
 "fmt"
"github.com/jinzhu/gorm"
 _ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
)

var DB *gorm.DB

func ConnectDataBase() {

    connection := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", "localhost", "5432", "postgres", "postgres", "postgres")


    database, err := gorm.Open("postgres", connection)
  
    if err != nil {
      panic("Failed to connect to database!")
    }
  
    database.AutoMigrate(&URLCollection{})

  
    DB = database
  }


