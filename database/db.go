package database 


import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
  "goql/models"
) 

type createdb struct {
  db *gorm.DB  
}



var db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

func main() {
  if err != nil {
    panic("Error occured, database")
  }
}

var Database *createdb = &createdb{db: db}

func (r *createdb) CreateUser(user *models.User) (*models.User, error) {
  usr := &user
  result := r.db.Create(usr)
  if result.Error != nil {
    return user, nil 
  }
  return user, nil 
}