package models
import (
	"github.com/dgrijalva/jwt-go"
	u "src/utils"
	"strings"
)
type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Account struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token";sql:"-"`
}

func (account *Account) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Password must have length longer than 6"), false
	}
	// Email must be unique
	temp := &Account{}

	// Check for errors and duplicate emails

	return u.Message(true, "Requirement passed"), true
}
//
//func (account *Account) Create() map[string] interface{} {
//
//}