DROP USER "crud-6";
CREATE USER "crud-6" WITH PASSWORD '12345' SUPERUSER;

DROP DATABASE "crud-6-db";
CREATE DATABASE "crud-6-db" OWNER "crud-6";

\connect crud-6-db

DROP TABLE users;
CREATE TABLE users (
    id INT,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    dob  TIMESTAMP,
    location VARCHAR(50),
    cellphone_number VARCHAR(50),
    email VARCHAR(50),
    password VARCHAR(50),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
ALTER TABLE persons OWNER TO "crud-6";

CREATE TABLE books (
    id INT,
    tittle VARCHAR(50),
    author VARCHAR(50),
    pages INT,
    quantity INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE users_books (
    id_user INTEGER,
    id_book INTEGER
);

