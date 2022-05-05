create database meetings;
use meetings;

create table meetings(
    id int auto_increment primary key,`resource` varchar(255) not null,
    user_id int not null,
    topic varchar(255) not null,
    application_id BIGINT not null,
    attempts int not null,
    `sent` timestamp not null,
    received timestamp not null
);
