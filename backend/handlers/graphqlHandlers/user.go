package graphqlHandlers

import (
	"Knoxiaes/fairesults/database"
	"Knoxiaes/fairesults/database/users"
	"Knoxiaes/fairesults/graph/model"
	"Knoxiaes/fairesults/helpers"
	"log"
)

func User(username string, numberOfResults int) (*model.User, error) {
	log.Println(numberOfResults)
	res := database.DB.QueryRow("select user_id,username,email,firstname,lastname,birthday from users where username=?", username)
	var tmpUser users.User
	err := res.Scan(&tmpUser.UserID, &tmpUser.Username, &tmpUser.Email, &tmpUser.Firstname, &tmpUser.Lastname, &tmpUser.Birthday)
	if err != nil {
		log.Println(err)
		return nil, helpers.CustomError{Message: err.Error(), Code: 0}
	}
	graphqlResults, err := Results(tmpUser.UserID, numberOfResults)
	if err != nil {
		log.Println(err)
		return nil, helpers.CustomError{Message: err.Error(), Code: 0}
	}
	return &model.User{
		UserID:    tmpUser.UserID,
		Username:  tmpUser.Username,
		Email:     tmpUser.Email,
		Firstname: tmpUser.Firstname,
		Lastname:  tmpUser.Lastname,
		Birthday:  tmpUser.Birthday,
		Results:   graphqlResults,
	}, nil
}
