package users

import (
	"Knoxianes/fairecords/db/users/query"
	"Knoxianes/fairecords/db/users/schema"
	"Knoxianes/fairecords/helpers"
)

func GetUserByUsername(username string) (schema.User, error) {
	var user schema.User

	db, err := helpers.DBConnect()
	if err != nil {
		return user, err
	}
    defer db.Close()

	if err := db.QueryRow(query.GetUserByUsername, username).Scan(&user.Username,&user.Password,&user.Email,&user.Token,&user.Verification_token,&user.Verified); err != nil {
		return user, err
	}

	return user, nil
}

func GetUserByEmail(email string) (schema.User, error) {
	var user schema.User

	db, err := helpers.DBConnect()
	if err != nil {
		return user, err
	}
    defer db.Close()

	if err := db.QueryRow(query.GetUserByEmail, email).Scan(&user.Username,&user.Password,&user.Email,&user.Token,&user.Verification_token,&user.Verified); err != nil {
		return user, err
	}

	return user, nil
}


func InsertUser(user schema.User) error {

	db, err := helpers.DBConnect()
	if err != nil {
		return err
	}
    defer db.Close()

	if _, err := db.Exec(query.InsertUser, user.Username, user.Password, user.Email, user.Token.String, user.Verification_token.String); err != nil {
		return err
	}
	return nil
}

func DeleteUser(username string) error{
    
    db, err := helpers.DBConnect()
    if err != nil{
        return err
    }
    defer db.Close()

    if _,err := db.Exec(query.DeleteUser,username); err !=nil {
        return err
    }

    return nil
}

func UpdateUserPassword(username string, password string) error {
    db, err := helpers.DBConnect()
    if err != nil{
        return err
    }
    defer db.Close()

    if _,err := db.Exec(query.UpdateUserPassword,password,username); err !=nil {
        return err
    }

    return nil
}
func UpdateUserEmail(username string, email string) error {
    db, err := helpers.DBConnect()
    if err != nil{
        return err
    }
    defer db.Close()

    if _,err := db.Exec(query.UpdateUserEmail,email,username); err !=nil {
        return err
    }

    return nil
}
func UpdateUserToken(username string, token string) error {
    db, err := helpers.DBConnect()
    if err != nil{
        return err
    }
    defer db.Close()

    if _,err := db.Exec(query.UpdateUserToken,token,username); err !=nil {
        return err
    }

    return nil
}
func UpdateUserVerificationToken(username string, verification_token string) error {
    db, err := helpers.DBConnect()
    if err != nil{
        return err
    }
    defer db.Close()

    if _,err := db.Exec(query.UpdateUserVerificationToken,verification_token,username); err !=nil {
        return err
    }

    return nil
}
func UpdateUserVerified(username string, verified bool) error {
    db, err := helpers.DBConnect()
    if err != nil{
        return err
    }
    defer db.Close()

    if _,err := db.Exec(query.UpdateUserVerified,verified,username); err !=nil {
        return err
    }

    return nil
}
