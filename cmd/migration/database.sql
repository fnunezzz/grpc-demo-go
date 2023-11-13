create schema if not exists demo_grpc;

create table if not exists demo_grpc.users (
    id uuid,
    username varchar(50) not null unique,
    email varchar(200) not null,
    password text not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    primary key (id)
);