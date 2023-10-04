package logic

import (
	"context"
	"log"
	"rto/database"
	"rto/graph/model"
)

func SignUpUser(ctx context.Context, input *model.SignUpRequest) (*model.SignUpResponse, error) {
	var response model.SignUpResponse
	// var msg string
	// msg = "User Created successfully"

	pass, err := Encrypt(*input.Password)
	if err != nil {
		return nil, err
	}

	err2 := database.Upsertuser(ctx, input, pass)
	if err2 != nil {
		log.Println(err.Error(), "<- Error")
	}

	passValue, err3 := Decrypt(pass)
	if err3 != nil {
		return nil, err3
	}

	response.Message = &passValue
	return &response, nil
}
