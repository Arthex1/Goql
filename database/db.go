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


type BioDB struct {
	URLCode     string   `json:"url_code"`
	Links       *LinksDB  `gorm:"foreignKey:ID"`
	UserID      string   `json:"user_id"`
	Description string   `json:"description"`
	Skills       []*models.Skill `gorm:"foreignKey:UserID;references:ID"` 
}

type LinksDB struct {
  ID string
	Portfolio string `json:"portfolio"`
	Github    string `json:"github"`
	Youtube   string `json:"youtube"`
	Twitter   string `json:"twitter"`
}

// type SkillDB struct {
//   ID string 
// 	Name string `json:"name"`
// }

type UserDB struct {
  Name string 
  ID string
  Bio *BioDB `gorm:"foreignKey:URLCode"`
}


var db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

func run() (success bool) {
  if err != nil {
    panic("Error occured, database")
  }
  log.Println("test")
  err := db.AutoMigrate(&UserDB{})
  if err != nil {
    log.Println(err) 
    return false 
  }
  
  return true 
}

var p any = run() 
var Database *createdb = &createdb{db: db}

func (r *createdb) CreateUser(user *models.User) (*models.User, error) {
  var usr *UserDB = User_to_user(user)
  result := r.db.Create(usr)
  if result.Error != nil {
    return user, nil 
  }
  return user, nil 
}

func (r *createdb) GetUsers(limit int) ([]*models.User, error) {
  var users []*UserDB 
  db.Limit(limit + 1).Find(&users)
  var converted []*models.User 
  for i := range users {
    converted = append(converted, User_to_userx(users[i]))
  } 
  return converted, nil 
   
}

func User_to_user(user *models.User) (userx *UserDB) {
  return &UserDB{
    ID: user.ID,
    Name: user.Name,
    Bio: &BioDB{
      UserID: user.ID,
      URLCode: user.Bio.URLCode,
      Description: user.Bio.Description,
      Links: &LinksDB{
        Github: user.Bio.Links.Github,
        Youtube: user.Bio.Links.Youtube,
        Portfolio: user.Bio.Links.Portfolio,
        Twitter: user.Bio.Links.Twitter,
      },
      Skills: user.Bio.Skils,
    },
  }
}

func User_to_userx(user *UserDB) (userx *models.User) {
  return &models.User{
    ID: user.ID,
    Name: user.Name,
    Bio: &models.Bio{
      UserID: user.ID,
      URLCode: user.Bio.URLCode,
      Description: user.Bio.Description,
      Links: &models.Links{
        Github: user.Bio.Links.Github,
        Youtube: user.Bio.Links.Youtube,
        Portfolio: user.Bio.Links.Portfolio,
        Twitter: user.Bio.Links.Twitter,
      },
      Skils: user.Bio.Skills,
    },
  }
}