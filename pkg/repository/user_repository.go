package repository

import (
	"context"

	"github.com/Prameesh-P/chat-api-cleancode/pkg/helper"
	"github.com/Prameesh-P/chat-api-cleancode/pkg/response"
	"github.com/Prameesh-P/chat-api-cleancode/pkg/models"
	"gorm.io/gorm"
)
type UserRepoInterface interface {
	UserSignup(ctx context.Context, user helper.UserHelper) (response.User, error)
	UserLogin(ctx context.Context,email string)(models.Users,error)
}

type UserDatabase struct{
	DB *gorm.DB
}

func NewUserRepo(DB *gorm.DB)UserRepoInterface{
	return &UserDatabase{DB: DB}
}

func (u *UserDatabase)UserSignup(ctx context.Context, user helper.UserHelper) (response.User, error){
	var users response.User
	insertQuery := `INSERT INTO users (name,email,mobile,password)VALUES($1,$2,$3,$4) 
					RETURNING id,name,email,mobile`
	err:=u.DB.Raw(insertQuery,user.Name,user.Email,user.Mobile,user.Password).Scan(&users).Error
	return users,err
}
func (u *UserDatabase)UserLogin(ctx context.Context,email string)(models.Users,error){
	var users models.Users
	err := u.DB.Raw("SELECT * FROM users WHERE email=?", email).Scan(&users).Error
	return users, err
}