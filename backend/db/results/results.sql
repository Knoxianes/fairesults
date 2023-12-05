CREATE table results(
    result_id integer primary key,
    username text not null,
    date integer not null,
    competition text not null,
    category integer not null,
    number_competitors integer not null,
    place integer not null,
    competition_rank real not null,
    mass_coefficient real not null,
    medal integer,
    record integer
);
