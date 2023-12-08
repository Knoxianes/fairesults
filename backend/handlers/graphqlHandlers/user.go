package graphqlHandlers

import (
	"Knoxiaes/fairesults/database"
	"Knoxiaes/fairesults/database/users"
	"Knoxiaes/fairesults/graph/model"
	"log"
	"strconv"
)

func User(userID int, numberOfResults int) (*model.User, error) {
	log.Println(numberOfResults)
	res := database.DB.QueryRow("select user_id,username,email,firstname,lastname,birthday from users where user_id=?", userID)
	var tmpUser users.User
	err := res.Scan(&tmpUser.UserID, &tmpUser.Username, &tmpUser.Email, &tmpUser.Firstname, &tmpUser.Lastname, &tmpUser.Birthday)
	if err != nil {
		log.Println(err)
		return &model.User{}, err
	}
	graphqlUserID := strconv.Itoa(tmpUser.UserID)
	graphqlResults, err := Results(userID, 0)
	if err != nil {
		log.Println(err)
		return &model.User{}, err
	}
	return &model.User{
		UserID:    graphqlUserID,
		Username:  tmpUser.Username,
		Email:     tmpUser.Email,
		Firstname: tmpUser.Firstname,
		Lastname:  tmpUser.Lastname,
		Birthday:  tmpUser.Birthday,
		Results:   graphqlResults,
	}, nil
}
