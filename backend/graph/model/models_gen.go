// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewResult struct {
	CompetitionName      string  `json:"competitionName"`
	Category             int     `json:"category"`
	NumberOfCompetitiors int     `json:"numberOfCompetitiors"`
	Place                int     `json:"place"`
	CompetitionRank      float64 `json:"competitionRank"`
	Date                 int     `json:"date"`
	Medal                int     `json:"medal"`
	Record               int     `json:"record"`
}

type NewUser struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Birthday  int    `json:"birthday"`
}

type Result struct {
	ResultID            int     `json:"resultID"`
	CompetitionName     string  `json:"competitionName"`
	Category            int     `json:"category"`
	NumberOfCompetitors int     `json:"numberOfCompetitors"`
	Place               int     `json:"place"`
	CompetitionRank     float64 `json:"competitionRank"`
	Date                int     `json:"date"`
	MassCoefficient     float64 `json:"massCoefficient"`
	Medal               int     `json:"medal"`
	Record              int     `json:"record"`
	Points              float64 `json:"points"`
}

type UpdatePassword struct {
	OldPassword *string `json:"oldPassword,omitempty"`
	NewPassword string  `json:"newPassword"`
}

type UpdatedResult struct {
	ResultID            int     `json:"resultID"`
	CompetitionName     string  `json:"competitionName"`
	Category            int     `json:"category"`
	NumberOfCompetitors int     `json:"numberOfCompetitors"`
	Place               int     `json:"place"`
	CompetitionRank     float64 `json:"competitionRank"`
	Date                int     `json:"date"`
	MassCoefficient     float64 `json:"massCoefficient"`
	Medal               int     `json:"medal"`
	Record              int     `json:"record"`
	Points              float64 `json:"points"`
}

type User struct {
	UserID    int       `json:"userID"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Birthday  int       `json:"birthday"`
	Results   []*Result `json:"results"`
}
