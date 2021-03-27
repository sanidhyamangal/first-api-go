// Author: Sanidhya Mangal
// github: sanidhyamangal
package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sanidhya/firstapi/config"
	"github.com/sanidhya/firstapi/models"
	"github.com/sanidhya/firstapi/routes"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("status:", err)
	}

	defer config.DB.Close()

	config.DB.AutoMigrate(&models.User{})

	r := routes.SetUpRouter()

	r.Run()
}
