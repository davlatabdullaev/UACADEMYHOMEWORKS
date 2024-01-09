create table drivers (
   id uuid primary key,
   full_name text,
   phone text not null 
);
create table cars (
  id uuid primary key,
  model text not null,
  brand text not null,
  year int,
);