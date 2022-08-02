create database todo_list;
use todo_list;

-- Mysql
create table todo_items
(
    id         int auto_increment primary key ,
    title      varchar(100),
    content    varchar(255),
    status     enum ('doing', 'review','done'),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Postgresql
CREATE TYPE status_items_enum AS ENUM ('doing', 'review','done');
create table todo_items (
    id         serial primary key ,
    title      varchar(100),
    content    varchar(255),
    status     status_items_enum,
    created_at timestamptz,
    updated_at timestamptz
);
