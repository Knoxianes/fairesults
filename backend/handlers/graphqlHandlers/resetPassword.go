package graphqlHandlers

import (
	"Knoxiaes/fairesults/database"
	"Knoxiaes/fairesults/helpers"
	"database/sql"
	"log"
)

func ResetPassword(username string, newPassword string)(bool,error){
	row := database.DB.QueryRow("select verified from users where username = ?;",username)
	var verified int

	err := row.Scan(&verified)
	if err != nil{
		if err == sql.ErrNoRows{
			log.Println(err)
			return false,helpers.CustomError{Message:"User doesnt exists",Code:3}
		}else{
			log.Println(err)
			return false,helpers.CustomError{Message:err.Error(),Code:0}
		}
	}

	if verified == 0{
		return false, helpers.CustomError{Message:"User not verified", Code:4}
	}

	hashedPassword, err := helpers.HashPassword(newPassword)
	if err != nil{
		log.Println(err)
		return false, helpers.CustomError{Message:err.Error(),Code:0}
	}

	_, err = database.DB.Query("update users set password=? where username=?;",hashedPassword,username)
	if err != nil{
		log.Println(err)
		return false, helpers.CustomError{Message:err.Error(),Code:0}
	}
	
	return true,nil
}
