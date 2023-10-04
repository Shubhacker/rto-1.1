package database

import (
	"context"
	"log"
	"rto/graph/model"

	"github.com/google/uuid"
)

func Upsertuser(ctx context.Context, input *model.SignUpRequest, pass string) error {
	// var DBResponse model.SignUpResponse
	var inputArgs []interface{}

	id := uuid.New()

	query := `insert into public.master_user (userid, username, userfirstname, userlastname, createdat, email, isactive, "password" ) values(
		$1, $2, $3, $4, current_timestamp, $5, true, $6) returning userid;`

	inputArgs = append(inputArgs, id)
	inputArgs = append(inputArgs, input.Username)
	inputArgs = append(inputArgs, input.FirstName)
	inputArgs = append(inputArgs, input.LastName)
	inputArgs = append(inputArgs, input.Email)
	inputArgs = append(inputArgs, pass)

	err2 := DB.QueryRow(query, inputArgs...).Scan(&id)
	if err2 != nil {
		log.Println("Error from DB")
	}

	return nil
}
