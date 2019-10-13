package models

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Token - jwt token struct
type Token struct {
	UserID uint
	jwt.StandardClaims
}

// User - user struct
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token" sql:"-"`
}

// Create - create a user
func (user *User) Create() map[string]interface{} {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	DB.Create(user)

	if user.ID <= 0 {
		return map[string]interface{}{"message": "Failed to create user"}
	}

	user.Password = "" //delete password

	res := map[string]interface{}{}
	res["user"] = user
	return res
}

// Login - authenticate a user
func Login(username, password string) map[string]interface{} {
	user := &User{}

	err := DB.Table("users").Where("username = ?", username).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return map[string]interface{}{"message": "User does not exist"}
		}
		return map[string]interface{}{"message": "Failed to connect to db"}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return map[string]interface{}{"message": "Invalid username/password"}
	}

	user.Password = ""

	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_secret")))
	user.Token = tokenString

	res := map[string]interface{}{}
	res["user"] = user
	return res
}

// GetUsers -
func GetUsers() []User {
	users := []User{}
	DB.Table("users").Find(&users)
	// users, err := DB.Table("users").Find(&users).Error
	if err != nil {
		return nil
	}
	return users
}
