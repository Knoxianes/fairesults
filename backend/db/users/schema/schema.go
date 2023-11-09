package schema

import "database/sql"

type User struct{
    Username string
    Password string
    Email string
    Token sql.NullString
    Verification_token sql.NullString
    Verified sql.NullInt64
}
