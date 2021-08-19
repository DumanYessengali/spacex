create table auths
(
    id        serial primary key,
    user_id   int          not null,
    auth_uuid varchar(255) not null
);

create table users
(
    id serial primary key,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password varchar(255) not null,
    email varchar(255) not null unique,
    role varchar(255)
);