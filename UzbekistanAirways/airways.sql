create table tickets (
     id                 uuid primary key not null,
     from_city          text,
     to_city            text,
     flight_date        varchar(10)
);
create table passengers (
    id                  uuid primary key not null,
    first_name          varchar(20),
    last_name           varchar(20),
    email               varchar(30),
    phone               varchar(13),
    ticket_id           uuid references tickets(id)
);