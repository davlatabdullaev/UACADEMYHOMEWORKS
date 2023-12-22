create table products (
id uuid primary key not null,
name varchar(50) not null,
description varchar(50),
price numeric,
 created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);
select id, name, description, price, created_at, updated_at
		from products
		where name ilike (%a%) AND price > 5000
select id, name, description, price, created_at, updated_at
from products
where (name LIKE '%burger%' OR description LIKE '%burger%') and created_at < updated_at;
       