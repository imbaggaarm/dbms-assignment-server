package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"
	u "src/utils"
	"strings"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Student struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string `gorm:"column:uname"`
	Password  string `gorm:"column:pass"`
	Location  string
	ImageUrl  string
	Token     string `sql:"-"`
}

func (Student) TableName() string {
	return "student"
}

func (account *Student) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Password must have length longer than 6"), false
	}
	// Email must be unique
	temp := &Student{}

	// Check for errors and duplicate emails
	err := GetDB().Table("student").Where("uname = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}

	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(true, "Requirement passed"), true
}

func (account *Student) Create() map[string]interface{} {
	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	// Create new JWT Token for the newly registered account
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" // Delete password
	response := u.Message(true, "Student has been created")
	response["data"] = account
	return response
}

func DeleteAccount(user uint) map[string]interface{} {
	account := Student{ID: user}
	GetDB().Delete(&account)
	// Call delete
	response := u.Message(true, "Student has been deleted")
	return response
}

func GetProfile(user uint) map[string]interface{} {
	account := Student{ID: user}
	GetDB().First(&account)
	response := u.Message(true, "")
	account.Password = ""
	response["data"] = account
	return response
}

func (account *Student) Update(firstName, lastName string) map[string]interface{} {
	GetDB().First(&account)
	account.FirstName = firstName
	account.LastName = lastName
	GetDB().Save(&account)
	response := u.Message(true, "Student has been updated")
	account.Password = ""
	response["data"] = account
	return response
}

func Login(email, password string) map[string]interface{} {
	account := &Student{}
	err := GetDB().Table("student").Where("uname = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again")
	}
	//Worked! Logged In
	account.Password = ""
	//Create JWT token
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString // Store the toke in the response

	resp := u.Message(true, "Logged In")
	resp["data"] = account
	return resp
}
