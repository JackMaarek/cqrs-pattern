package conf

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

// InitializeDb connects to database
func InitializeDb(user string, password string, host string, name string, port int) {
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, name)
	var tmpDb, err = gorm.Open("mysql", dbURL)
	if err != nil {
		fmt.Printf("Cannot connect to database")
		log.Fatal("error:", err)
		return
	}
	fmt.Printf("We are connected to database \n")
	DB = tmpDb
}
