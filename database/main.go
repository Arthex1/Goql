package database

import (
	"fmt"
	"goql/graph/model"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string    `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name"`
	Bio      Bio       `gorm:"foreignKey:ID" json:"bio"`
	Badges   Badges    `gorm:"foreignKey:id"`
	Links    Links     `gorm:"foreignKey:id"`
	Projects []Project `gorm:"foreignKey:id"`
}

type Bio struct {
	ID    string `json:"id"`
	Text  string `json:"text"`
	Email string `json:"email"`
}

type Links struct {
	ID        string
	Youtube   string
	Twitter   string
	LinkedIN  string
	Portfolio string
	Github    string
}

type Project struct {
	ID      string
	Link    string
	Picture string
	Name    string
}

type Badges struct {
	ID        string
	Developer bool
	PlusUser  bool
}

type database struct {
	db *gorm.DB
}


func Convert_projects_to_db(n []*model.ProjectInput) []Project {
	var d []Project
	for _, v := range n {
		m := append(d, Project{
			Name: v.Name,
			Link: v.Link,
			Picture: v.Picture,
		})
		d = m  	

	}
	return d
}

func Convert_projects_to_query(n []Project) []*model.Project {
	var d []*model.Project 
	for _, v := range n {
		m := append(d, &model.Project{
			Name: v.Name,
			Link: v.Link,
			Picture: v.Picture,
			ID: v.ID,
		})
		d = m 
	}
	return d 
}

func create_db() *database {
	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	if err != nil {
		panic("Database Had a error")
	}
	db.AutoMigrate(&User{}, &Bio{}, &Badges{}, &Links{}, &Project{})
	return &database{db: db}

}

var DB *database = create_db()

func (s *database) Create_user(name string, bio_text string, email string, developer bool, plus bool, youtube string, twitter string, linkedin string, portfolio string, github string, Project []Project) (User, error) {
	id := uuid.New().String()
	usr := s.db.Create(&User{Name: name, Bio: Bio{Text: bio_text, Email: email}, ID: id, Badges: Badges{Developer: developer, PlusUser: plus}, Links: Links{Youtube: youtube, Twitter: twitter, LinkedIN: linkedin, Portfolio: portfolio, Github: github}, Projects: Project})
	to_ret := new(User)

	if usr.Error != nil {
		return *to_ret, fmt.Errorf(usr.Error.Error())

	}

	usr_search := s.db.Model(&User{}).Where("id = ?", id).Preload("Bio").Preload("Badges").Preload("Links").Preload("Projects").First(&to_ret)
	
	if usr_search.Error != nil {
		return *to_ret, fmt.Errorf(usr_search.Error.Error())
	}

	return *to_ret, nil

}
func (s *database) Get_user() {
	

}