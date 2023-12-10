package graphqlHandlers

import (
	"Knoxiaes/fairesults/database"
	"Knoxiaes/fairesults/graph/model"
	"Knoxiaes/fairesults/helpers"
	"log"
)

func UpdatePassword(username string, input model.UpdatePassword) (bool, error) {
	row := database.DB.QueryRow("select password from users where username = ?;", username)
	var password string

	err := row.Scan(&password)
	if err != nil {
		log.Println(err)
		return false, helpers.CustomError{Message: err.Error(), Code: 0}
	}

	if helpers.CheckPasswordHash(*input.OldPassword,password) {
		return false, helpers.CustomError{Message: "Wrong old password", Code: 5}
	}

	hashedPassword, err := helpers.HashPassword(input.NewPassword)
	if err != nil {
		log.Println(err)
		return false, helpers.CustomError{Message: err.Error(), Code: 0}
	}

	_, err = database.DB.Query("update users set password=? where username=?;", hashedPassword, username)
	if err != nil {
		log.Println(err)
		return false, helpers.CustomError{Message: err.Error(), Code: 0}
	}

	return true, nil
}
