// Author: Sanidhya Mangal
// github: sanidhyamangal
package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sanidhya/firstapi/config"
)

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (b *User) TableName() string {
	return "user"
}

func GetAllUsers(user *[]User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func GetUserByID(user *User, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func UpdateUser(user *User, id string) (err error) {
	fmt.Println(user)
	config.DB.Save(user)
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(user *User, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(user)
	return nil
}
