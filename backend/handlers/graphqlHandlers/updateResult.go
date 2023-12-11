package graphqlHandlers

import (
	"Knoxiaes/fairesults/database"
	"Knoxiaes/fairesults/database/results"
	"Knoxiaes/fairesults/graph/model"
	"Knoxiaes/fairesults/helpers"
	"log"
)

func UpdateResult(username string, input model.UpdatedResult) (bool, error) {
	row := database.DB.QueryRow("select user_id from users where username =?;", username)

	var userID int
	err := row.Scan(&userID)
	if err != nil {
		log.Println(err)
		return false, helpers.CustomError{Message: err.Error(), Code: 0}
	}

	row = database.DB.QueryRow("select user_id from users where result_id = ?;", input.ResultID)
	var resultUserID int
	err = row.Scan(&resultUserID)
	if err != nil {
		log.Println(err)
		return false, helpers.CustomError{Message: "Result doesnt exists", Code: 7}
	}

	if resultUserID != userID {
		return false, helpers.CustomError{Message: "Warning unauthorized access", Code: -1}
	}

	result := results.Result{
		CompetitionName:     input.CompetitionName,
		Category:            input.Category,
		NumberOfCompetitors: input.NumberOfCompetitors,
		Place:               input.Place,
		CompetitionRank:     input.CompetitionRank,
		Date:                int64(input.Date),
		Medal:               input.Medal,
		Record:              input.Record,
	}
	result.CalculatePoints()

	_, err = database.DB.Query(`update results set 
								competition_name = ?, category = ?,  number_competitors = ?,
								place = ?, competition_rank = ?, date = ?, mass_coefficient = ?, 
								medal = ?, record = ?, points = ? where result_id = ?;`,
	result.CompetitionName, result.Category, result.NumberOfCompetitors, result.Place,
	result.CompetitionRank, result.Date, result.MassCoefficinet, result.Medal, result.Record, result.Points)

	if err != nil {
		log.Println(err)
		return false, helpers.CustomError{Message: err.Error(), Code: 0}
	}

	return true, nil
}
