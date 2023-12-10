package graphqlHandlers

import (
	"Knoxiaes/fairesults/database"
	"Knoxiaes/fairesults/database/results"
	"Knoxiaes/fairesults/graph/model"
	"Knoxiaes/fairesults/helpers"
	"log"
)

func CreateResult(username string, input model.NewResult) (*model.Result, error) {
	result := results.Result{
		CompetitionName:     input.CompetitionName,
		Category:            input.Category,
		NumberOfCompetitors: input.NumberOfCompetitiors,
		Place:               input.Place,
		CompetitionRank:     input.CompetitionRank,
		Date:                int64(input.Date),
		Medal:               input.Medal,
		Record:              input.Record,
	}
	result.CalculatePoints()
	userID, err := database.GetUserIDFromUsername(username)
	if err != nil {
		log.Println(err)
		return nil, helpers.CustomError{Message: err.Error(), Code: 0}
	}
	_, err = database.DB.Query(`insert into results
		(user_id,competition_name,category,number_competitors,competition_rank,place,date,mass_coefficient,medal,record,points)
		values(?,?,?,?,?,?,?,?,?,?,?)`,
		userID, result.CompetitionName, result.Category, result.NumberOfCompetitors, result.CompetitionRank, result.Place, result.Date, result.MassCoefficinet, result.Medal, result.Record, result.Points)
	if err != nil {
		log.Println(err)
		return nil, helpers.CustomError{Message: err.Error(), Code: 0}
	}
	resultID, err := database.GetLastInsertedIDFromResults()
	if err != nil {
		log.Println(err)
		return nil, helpers.CustomError{Message: err.Error(), Code: 0}
	}


	return &model.Result{
		ResultID:            resultID,
		CompetitionName:     result.CompetitionName,
		Category:            result.Category,
		NumberOfCompetitors: result.NumberOfCompetitors,
		Place:               result.Place,
		CompetitionRank:     result.CompetitionRank,
		Date:                input.Date,
		MassCoefficient:     result.MassCoefficinet,
		Medal:               result.Medal,
		Record:              result.Record,
		Points:              result.Points,
	}, nil
}
