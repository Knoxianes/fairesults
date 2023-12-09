package graphqlHandlers

import (
	"Knoxiaes/fairesults/database"
	"Knoxiaes/fairesults/graph/model"
	"Knoxiaes/fairesults/helpers"
	"database/sql"
	"log"
	"strconv"
	"time"
)

func CreateUser(input model.NewUser) (string, error) {
	row := database.DB.QueryRow("select user_id from users where username=?;", input.Username)
	err := row.Scan()
	if err != sql.ErrNoRows {
		log.Println(row.Err())
		return "-1", helpers.CustomError{Message: "Username already exists", Code: 500}
	}


	row = database.DB.QueryRow("select user_id from users where email=?;", input.Email)
	err = row.Scan()
	if err != sql.ErrNoRows {
		log.Println(row.Err())
		return "-1", helpers.CustomError{Message: "Email already exists", Code: 500}
	}


	emailVerificationToken, err := helpers.GenerateToken(input.Username, time.Hour)
	if err != nil {
		log.Println(err)
		return "-1", err
	}


	input.Password, err = helpers.HashPassword(input.Password)
	if err != nil {
		log.Println(err)
		return "-1", err
	}


	err = helpers.SendVerificationMailEmail(emailVerificationToken, input.Email)
	if err != nil {
		log.Println(err)
		return "-1", err
	}


	_, err = database.DB.Query("insert into users(username,password,email,firstname,lastname,birthday,verified) values (?,?,?,?,?,?,0);",
		input.Username, input.Password, input.Email, input.Firstname, input.Lastname, input.Birthday)
	if err != nil {
		log.Println(err)
		return "-1", err
	}


	id, err := database.GetLastInsertedIDFromUsers()
	if err != nil {
		log.Println(err)
		return "-1", err
	}



	return strconv.Itoa(id), nil
}
