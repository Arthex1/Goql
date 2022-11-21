package database

import (

	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
  "github.com/google/uuid"
)

type User struct {
  gorm.Model
  ID string `gorm:"primaryKey"`
  Name string 
  Bio Bio `gorm:"foreignKey:ID"`

}



type Bio struct {
  ID string 
  Text string
  Email string 
  
}






type database struct {
  db *gorm.DB

}

func create_db() *database {
  db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
  if err != nil {
    panic("Database Had a error")
  }
  db.AutoMigrate(&User{}, &Bio{}) 
  return &database{db: db}
  

}

var DB *database = create_db()

func (s *database) Create_user(name string, bio_text string, email string) (User, error) {
  id := uuid.New().String()
  usr := s.db.Create(&User{Name: name, Bio: Bio{Text: bio_text, Email: email}, ID: id })
  to_ret := new(User)


  if usr.Error != nil {
    return *to_ret, fmt.Errorf(usr.Error.Error())

  }

  usr_search := s.db.Model(&User{}).Where("id = ?", id).Preload("Bio").First(&to_ret)

  if usr_search.Error != nil {
    return *to_ret, fmt.Errorf(usr_search.Error.Error()) 
  }

  return *to_ret, nil 
  





}