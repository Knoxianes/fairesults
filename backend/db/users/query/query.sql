-- name: GetUserByUsername
SELECT * FROM users
WHERE username=? LIMIT 1;

-- name: GetUserByEmail
SELECT * FROM users
WHERE email=? LIMIT 1;

-- name: InsertUser 
INSERT INTO users(
    username,password,email,
    token,verification_token,verified
) VALUES(
    ?,?,?,?,?,FALSE
)

-- name: UpdateUserPassword
UPDATE users 
set password=?
WHERE username=?;

-- name: UpdateUserEmail
UPDATE users
set email=?
WHERE username=?;

-- name: UpdateUserToken
UPDATE users
set token=?
WHERE username=?;

-- name: UpdateUserVerificationToken
UPDATE users
set verification_token=?
WHERE username=?;

-- name: UpdateUserVerified
UPDATE users
set verified=?
WHERE username=?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE username=?;
