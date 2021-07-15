package main

import (
	"Product1/config"
	"Product1/model"
	"Product1/rest"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	var err error
	config.DB, err = gorm.Open("mysql", config.DbConnectionString(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Error Connecting to Infix DB: ", err)
	}

	defer config.DB.Close()

	config.DB.AutoMigrate(&model.Inventory{})

	rest.SetupRouter()

}
