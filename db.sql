create database todolist;
use todolist;

create table todo_items
(
    id         int auto_increment primary key ,
    title      varchar(10),
    content    varchar(50),
    status     enum ('doing', 'review','done'),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
