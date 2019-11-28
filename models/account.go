package models

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/ottmartens/mentor-server/utils"
	"github.com/ottmartens/mentor-server/utils/enums"
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
	Email      string `json:"email"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Tagline    string `json:"tagline"`
	Degree     string `json:"degree"`
	Year       string `json:"year"`
	Token      string `json:"token";gorm:"-"`
	Role       string `json:"role"`
	GroupId    *uint  `json:"groupId"`
	ImageUrl   string `json:"imageUrl"`
	Bio        string `json:"bio"`
	IsVerified bool   `json:"isVerified"`
}

type AuthResponse struct {
	Name     string `json:"name"`
	ImageUrl string `json:"imageUrl"`
	Token    string `json:"token"`
	Role     string `json:"role"`
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

	if account.Role != enums.UserTypes.Mentee && account.Role != enums.UserTypes.Mentor {
		return utils.Message(
			false, fmt.Sprintf("Role must me either %s or %s", enums.UserTypes.Mentee, enums.UserTypes.Mentor)), false
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
	response["data"] = AuthResponse{
		Name:     account.Name,
		ImageUrl: account.ImageUrl,
		Token:    account.Token,
		Role:     account.Role,
	}
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

	resp["data"] = AuthResponse{
		Name:     account.Name,
		ImageUrl: account.ImageUrl,
		Token:    account.Token,
		Role:     account.Role,
	}
	return resp
}

func (account *Account) GetPublicInfo() AccountPublic {
	return AccountPublic{
		Name:     account.Name,
		Tagline:  account.Tagline,
		Degree:   account.Degree,
		Year:     account.Year,
		UserId:   account.ID,
		ImageUrl: account.ImageUrl,
		Bio:      account.Bio,
	}
}

func GetUser(userId uint, hidePassword bool) *Account {

	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", userId).First(acc)
	if acc.Email == "" {
		return nil
	}

	if hidePassword {
		acc.Password = ""
	}

	return acc
}

func (account *Account) SetGroupId(groupId uint) {
	fmt.Println(groupId)
	account.GroupId = &groupId
	GetDB().Save(account)
}
