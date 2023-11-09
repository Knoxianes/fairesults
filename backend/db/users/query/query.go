package query

var GetUserByUsername = `SELECT * FROM users 
                WHERE username = ? LIMIT 1;`

var GetUserByEmail = `SELECT * FROM users 
                WHERE email = ? LIMIT 1;`

var InsertUser = `INSERT INTO users
                    (username,password,email, token,verification_token,verified) 
                    VALUES(?,?,?,?,?,FALSE);`


var DeleteUser = `DELETE FROM users 
                    WHERE username = ?;`


var UpdateUserPassword =`UPDATE users
                        set password=?
                        WHERE username=?;`

var UpdateUserEmail=`UPDATE users
                        set email=?
                        WHERE username=?;`

var UpdateUserToken=`UPDATE users
                        set token=?
                        WHERE username=?;`

var UpdateUserVerificationToken=`UPDATE users
                        set verification_token=?
                        WHERE username=?;`

var UpdateUserVerified=`UPDATE users
                        set verified=?
                        WHERE username=?;`
