package models

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Claims - jwt token struct
type Claims struct {
	UserID uint
	jwt.StandardClaims
}

// User - user struct
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

// Create - create a user with hashed password
func (user *User) Create() map[string]interface{} {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	DB.Create(user)

	if user.ID <= 0 {
		return map[string]interface{}{"response": "Failed to create user"}
	}

	return map[string]interface{}{"response": "Successfully created a user"}
}

// Login - authenticate a user
func Login(username, password string) map[string]interface{} {
	user := &User{}

	err := DB.Table("users").Where("username = ?", username).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return map[string]interface{}{"response": "User does not exist"}
		}
		return map[string]interface{}{"response": "Failed to connect to db"}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return map[string]interface{}{"response": "Invalid username/password"}
	}

	secret := os.Getenv("TOKEN_SECRET")

	claims := &Claims{UserID: user.ID}
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return map[string]interface{}{"token": signedToken}
}

// GetUsers - supply a collection of users
func GetUsers() []User {
	users := []User{}
	DB.Table("users").Find(&users)
	if err != nil {
		return nil
	}
	return users
}
