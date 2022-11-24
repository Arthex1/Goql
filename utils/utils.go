package utils

import (
	"encoding/json"
	"goql/database"
	

	"golang.org/x/crypto/bcrypt"
)

func TypeConverter[R any](data any) (*R, error) {
    var result R
    b, err := json.Marshal(&data)
    if err != nil {
      return nil, err
    }
    err = json.Unmarshal(b, &result)
    if err != nil {
      return nil, err
    }
    return &result, err
}



func HashPassword(password string) (string, error) {
  pass, err := bcrypt.GenerateFromPassword([]byte(password), 0) 
  if err != nil {
    return "", err 
  }
  return string(pass), nil 
}

func VerifyPassword(password string, id string) (bool, error) {
	usr, err := database.DB.Get_user("id", id)
	if err != nil {
		return false, err 
	}
	hasherr := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password)) 
	if hasherr != nil {
		return false, hasherr
	}
  
	return true, nil 


}



