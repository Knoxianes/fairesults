package graphqlHandlers

import (
	"Knoxiaes/fairesults/database"
	"Knoxiaes/fairesults/graph/model"
	"Knoxiaes/fairesults/helpers"
	"log"
)

func UpdateUser(username string, input model.NewUser) (bool, error) {
	row := database.DB.QueryRow("select password from users where username = ?;", username)
	var userPassword string

	err := row.Scan(&userPassword)
	if err != nil {
		log.Println(err)
		return false, helpers.CustomError{Message: err.Error(), Code: 0}
	}

	if !helpers.CheckPasswordHash(input.Password, userPassword) {
		return false, helpers.CustomError{Message: "Wrong password", Code: 5}
	}

	_, err = database.DB.Query(`update users set email = ?, firstname = ?, lastname = ?, birthday = ?
				where username = ?;`, input.Email, input.Firstname, input.Lastname, input.Birthday, username)

	if err != nil {
		log.Println(err)
		return false, helpers.CustomError{Message: err.Error(), Code: 0}
	}

	return true, nil
}
