CREATE TABLE category (
    id serial primary key,
    name varchar(15),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);
CREATE TABLE product (
    id serial primary key,
    name varchar(15),
    price numeric(8),
    category_id int references category(id),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);
