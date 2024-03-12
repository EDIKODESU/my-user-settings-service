-- +migrate Up

create table users (
    id bigserial primary key,
    first_name text not null,
    second_name text not null,
    login text not null unique,
    mail text not null unique,
    password text not null
);

-- +migrate Down
drop table users;