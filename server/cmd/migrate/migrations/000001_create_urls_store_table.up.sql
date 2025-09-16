create table if not exists urls_store (
   id        integer primary key,
   title     varchar(100),
   url       varchar(100) not null,
   shortlink varchar(50) not null,
   clicks    integer not null
);