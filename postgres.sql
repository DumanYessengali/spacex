create table auths
(
    id        serial primary key,
    user_id   int          not null,
    auth_uuid varchar(255) not null
);

create table users
(
    id       serial primary key,
    name     varchar(255) not null,
    username varchar(255) not null unique,
    password varchar(255) not null,
    email    varchar(255) not null unique,
    role     varchar(255)
);

create table user_informations
(
    id                serial primary key,
    user_avatar_image varchar(255),
    user_city         varchar(255),
    user_age          int,
    user_id           int references users (id) on delete cascade not null
);


create table course
(
    id                 serial primary key,
    course_name        varchar(255) not null,
    course_description varchar(255) not null
);

create table user_courses
(
    id        serial primary key,
    course_id int references courses (id) on delete cascade not null,
    user_id   int references users (id) on delete cascade   not null
);

create table video_posts
(
    id             serial primary key,
    title          varchar(255)                           not null unique,
    title_type     varchar(255)                           not null,
    video_duration int                                    not null,
    description    varchar(255)                           not null,
    video_url      varchar(255)                           not null,
    created        timestamp with time zone default now() not null,
    updated        timestamp with time zone default now() not null
);

create table article_posts
(
    id                           serial primary key,
    title                        varchar(255)                           not null unique,
    title_type                   varchar(255)                           not null,
    duration                     int                                    not null,
    author_information_paragraph varchar(255),
    paragraph_name               varchar(255),
    description                  varchar(255)                           not null,
    author_name                  varchar(255),
    author_position              varchar(255),
    created                      timestamp with time zone default now() not null,
    updated                      timestamp with time zone default now() not null
);

create table post_connections
(
    id        serial primary key,
    post_id   int          not null,
    post_type varchar(255) not null
);


create table user_saved_posts
(
    id                 serial primary key,
    post_connection_id int references post_connections (id) on delete cascade not null,
    user_id            int references users (id) on delete cascade            not null
);