CREATE TABLE countries (
    id uuid primary key not null,
    name varchar(40) not null,
    currency varchar(10)
);

CREATE TABLE cities (
    id uuid primary key not null,
    name varchar(40),
    population int default 0,
    country_id uuid references countries(id)
);