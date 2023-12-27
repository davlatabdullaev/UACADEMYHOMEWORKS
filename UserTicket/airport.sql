CREATE TABLE ticket (
    id            serial primary key not null,
    from_city     varchar(15),
    to_city       varchar(15),
    flight_date   varchar(10)
);
CREATE TABLE passenger (
    id            serial primary key not null,
    first_name    varchar(15),
    last_name     varchar(15),
    email         varchar(20),
    phone         varchar(13),
    ticket_id     int references ticket(id)
);

