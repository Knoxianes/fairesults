CREATE TABLE users (
    username TEXT  PRIMARY KEY,
    password TEXT NOT NUll,
    email TEXT UNIQUE NOT NULL,
    token TEXT,
    verification_token TEXT,
    verified INTEGER
);
