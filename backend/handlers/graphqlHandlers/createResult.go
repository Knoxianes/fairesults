package graphqlHandlers

import (
	"Knoxiaes/fairesults/database"
	"Knoxiaes/fairesults/database/results"
	"Knoxiaes/fairesults/graph/model"
	"log"
	"strconv"
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
		return nil, err
	}
	_, err = database.DB.Query(`insert into results
		(user_id,competition_name,category,number_competitors,competition_rank,place,date,mass_coefficient,medal,record,points)
		values(?,?,?,?,?,?,?,?,?,?,?)`,
		userID, result.CompetitionName, result.Category, result.NumberOfCompetitors,result.CompetitionRank, result.Place, result.Date, result.MassCoefficinet, result.Medal, result.Record, result.Points)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resultID, err := database.GetLastInsertedIDFromResults()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	graphQLResultID := strconv.Itoa(resultID)

	return &model.Result{
		ResultID:            graphQLResultID,
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
