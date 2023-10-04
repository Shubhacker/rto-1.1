package database

import "rto/structs"

func FetchUserInfo(username string) (structs.UserInfo, error) {
	var user structs.UserInfo

	query := `select userid , username , userfirstname , userlastname , email, "password"  from public.master_user mu where username = $1`

	rows, err := DB.Query(query, username)
	if err != nil {
		return user, err
	}

	for rows.Next() {
		rows.Scan(
			&user.UserId,
			&user.UserName,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
		)
	}

	return user, nil
}
