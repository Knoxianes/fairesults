package graphqlHandlers

import (
	"Knoxiaes/fairesults/database"
	"Knoxiaes/fairesults/database/results"
	"Knoxiaes/fairesults/graph/model"
	"Knoxiaes/fairesults/helpers"
	"database/sql"
	"log"
)

func Results(userID int, numberOfResults int) ([]*model.Result, error) {
	var res *sql.Rows
	var err error
	if numberOfResults == 0 {
		res, err = database.DB.Query("select * from results where user_id=? order by points desc;", userID, numberOfResults)

	} else {

		res, err = database.DB.Query("select * from results where user_id=? order by points desc limit ?;", userID, numberOfResults)
	}
	if err != nil {
		log.Println(err)
		return nil, helpers.CustomError{Message: err.Error(), Code: 0}
	}
	var graphqlResults []*model.Result
	for res.Next() {
		var tmpResult results.Result
		err := res.Scan(&tmpResult.ResultID, &tmpResult.UserID, &tmpResult.CompetitionName, &tmpResult.Category, &tmpResult.NumberOfCompetitors, &tmpResult.Place, &tmpResult.CompetitionRank,
			&tmpResult.Date, &tmpResult.MassCoefficinet, &tmpResult.Medal, &tmpResult.Record, &tmpResult.Points)
		if err != nil {
			log.Println(err)
			return nil, helpers.CustomError{Message: err.Error(), Code: 0}
		}
		graphqlResults = append(graphqlResults, &model.Result{
			ResultID:            tmpResult.ResultID,
			CompetitionName:     tmpResult.CompetitionName,
			Category:            tmpResult.Category,
			NumberOfCompetitors: tmpResult.NumberOfCompetitors,
			Place:               tmpResult.Place,
			CompetitionRank:     tmpResult.CompetitionRank,
			Date:                int(tmpResult.Date),
			MassCoefficient:     tmpResult.MassCoefficinet,
			Medal:               tmpResult.Medal,
			Record:              tmpResult.Record,
			Points:              tmpResult.Points,
		})

	}
	if err := res.Err(); err != nil {
		log.Println(err)
		return nil, helpers.CustomError{Message: err.Error(), Code: 0}
	}
	return graphqlResults, nil
}
