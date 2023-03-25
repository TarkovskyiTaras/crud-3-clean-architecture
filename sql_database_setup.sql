DROP USER "crud-6";
CREATE USER "crud-6" WITH PASSWORD '12345' SUPERUSER;

DROP DATABASE "crud-6-db";
CREATE DATABASE "crud-6-db" OWNER "crud-6";

\connect crud-6-db

DROP TABLE users;
CREATE TABLE users (
    id VARCHAR(50),
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    dob  DATE,
    location VARCHAR(50),
    cellphone_number VARCHAR(50),
    email VARCHAR(50),
    password VARCHAR(50),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

ALTER TABLE persons OWNER TO "crud-6";

{"id" : 3, "first_name" : "1111111111111", "last_name" : "1111111111111", "dob" : "1990-12-28T00:00:00Z", "location" : "1111111111111", "cellphone_number" : "1111111111111"}
curl -i -X PUT -H "Content-Type: application/json" -d '{"id" : 1, "first_name" : "1111111111111", "last_name" : "1111111111111", "dob" : "1990-12-28T00:00:00Z", "location" : "1111111111111", "cellphone_number" : "1111111111111"}' "127.0.0.1:8080/user/update"