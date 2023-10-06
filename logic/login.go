package logic

import (
	"context"
	"rto/database"
	"rto/graph/model"
)

func Login(ctx context.Context, input *model.Login) (*model.LoginResponse, error) {
	var response model.LoginResponse
	var authenticated bool
	var user model.UserInformation
	var msg string
	encryptPass, _ := Encrypt(input.Password)

	userInfo, err := database.FetchUserInfo(input.Username)
	if err != nil {
		return nil, err
	}

	if userInfo.Password == encryptPass {
		authenticated = true
	}
	response.Authenticated = &authenticated

	if !authenticated {
		msg = "Wrong Password"
		response.Message = &msg
		return &response, nil
	}
	msg = "Logic Success"

	user.FirstName = &userInfo.FirstName
	user.LastName = &userInfo.LastName
	user.UserID = &userInfo.UserId
	user.Username = &userInfo.UserName
	JWT, err := GenerateToken(&userInfo)
	if err != nil {
		return &response, err
	}
	user.Jwt = JWT

	response.Message = &msg
	response.User = &user
	return &response, nil
}
