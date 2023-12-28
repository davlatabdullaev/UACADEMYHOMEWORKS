CREATE TABLE company (
    id serial primary key not null,
    company_name varchar(15),
    employees_count int
);
CREATE TABLE employees (
    id uuid primary key not null,
    employee_name varchar(20),
    phone varchar(13),
    company_id int references company(id),
    salary int,
    emloyee_type varchar(25)
);