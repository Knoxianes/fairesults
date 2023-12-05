CREATE table users(
    username text primary key,
    email text unique not null,
    password text not null,
    firstname text not null,
    lastname text not null,
    birthday integer not null,
    jwt_token text not null,
    jwt_verification_token text not null,
    verified integer not null,
    rlb real,
    competition_points blob
);
