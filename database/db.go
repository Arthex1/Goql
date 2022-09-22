package database

import (
	"goql/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
) 

type createdb struct {
  db *gorm.DB  
}

var db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

func run() (success bool) {
  if err != nil {
    panic("Error occured, database")
  }
  log.Println("test")
  db.AutoMigrate(&models.User{})
  return true 
}

var p any = run() 
var Database *createdb = &createdb{db: db}

func (r *createdb) CreateUser(user *models.User) (*models.User, error) {
  var usr *models.User = user
  result := r.db.Create(usr)
  if result.Error != nil {
    return user, nil 
  }
  return user, nil 
}

func (r *createdb) GetUsers(limit int) ([]*models.User, error) {
  var users []*models.User 
  db.Limit(limit + 1).Find(&users) 
  return users, nil 
   
}