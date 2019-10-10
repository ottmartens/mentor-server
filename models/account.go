package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/ottmartens/mentor-server/utils"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}
type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";gorm:"-"`
}

func generateTokenWithId(id uint) string {

	tk := &Token{UserId: id}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))

	return tokenString
}

func (account *Account) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return utils.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 8 {
		return utils.Message(false, "Password must be at least 8 characters"), false
	}

	temp := &Account{}

	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Connection error, please retry"), false
	}

	if temp.Email != "" {
		return utils.Message(false, "Email address already in use"), false
	}

	return utils.Message(false, "Requirement passed"), true
}

func (account *Account) Create() map[string]interface{} {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID <= 0 {
		return utils.Message(false, "Failed to create account, connection error")
	}

	account.Token = generateTokenWithId(account.ID)
	account.Password = ""

	response := utils.Message(true, "Account has been created")
	response["account"] = account
	return response
}

func Login(email, password string) map[string]interface{} {

	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "Cannot find account associated with this email address")
		}
		return utils.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return utils.Message(false, "Invalid login credentials. Please try again")
		}
		return utils.Message(false, "Error authenticating. Please retry")
	}

	account.Token = generateTokenWithId(account.ID)
	account.Password = ""

	resp := utils.Message(true, "Logged in")
	resp["account"] = account
	return resp
}

//func GetUser(u uint) *Account {
//
//	acc := &Account{}
//	GetDB().Table("accounts").Where("id = ?", u).First(acc)
//	if acc.Email == "" {
//		return nil
//	}
//
//	acc.Password = ""
//	return acc
//}
