
CREATE TABLE users(
        user_id integer primary key,
        username text unique not null,
        email text unique not null,
        password text not null,
        firstname text not null,
        lastname text not null,
        birthday integer not null,
        verified integer not null
);
