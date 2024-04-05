-- Utwórz bazę danych user_data
CREATE DATABASE user_data;

-- Użyj bazy danych user_data
\c

-- Utwórz tabelę role
CREATE TABLE IF NOT EXISTS userrole (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- Utwórz tabelę user
CREATE TABLE IF NOT EXISTS userinfo (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(256) NOT NULL,
    role_id INT NOT NULL,
    FOREIGN KEY (role_id) REFERENCES userrole (id)
);

-- Dodaj rekordy do tabeli role
INSERT INTO userrole (name) VALUES ('user'), ('admin');

-- Dodaj rekord do tabeli user
INSERT INTO userinfo (email, password, role_id) VALUES ('user@gmail.com', '1234', 1);

