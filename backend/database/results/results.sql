
CREATE TABLE results(
        result_id integer primary key,
        user_id integer not null,
        competition_name text not null,
        category integer not null,
        number_competitors integer not null,
        place integer not null,
        competition_rank real not null,
        date integer not null,
        mass_coefficient real not null,
        medal integer not null,
        record integer not null,
        points real not null
);
