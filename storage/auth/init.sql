CREATE DATABASE auth IF NOT EXISTS;
CREATE TABLE users IF NOT EXISTS(
    id int primary,
    username varchar,
    pwd varchar,
    email varchar
);
