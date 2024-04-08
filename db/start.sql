-- Utwórz bazę danych user_data
CREATE DATABASE user_data;

-- Użyj bazy danych user_data
\c

CREATE TABLE IF NOT EXISTS userinfo (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(256) NOT NULL,
    role INT,
    FOREIGN KEY (role) REFERENCES userrole(id)
);

CREATE TABLE IF NOT EXISTS userrole (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    text VARCHAR(1500) NOT NULL
);

CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    text VARCHAR(300) NOT NULL,
    id_post INT,
    FOREIGN KEY (id_post) REFERENCES posts(id)
);

CREATE TABLE IF NOT EXISTS page (
    id SERIAL PRIMARY KEY,
    post_list JSON NOT NULL
);

CREATE TABLE IF NOT EXISTS postreactions (
    id SERIAL PRIMARY KEY,
    target_id INT,
    post_id INT,
    stats JSON NOT NULL,
    FOREIGN KEY (post_id) REFERENCES posts(id)
);

CREATE TABLE IF NOT EXISTS comreactions (
    id SERIAL PRIMARY KEY,
    target_id INT,
    comment_id INT,
    stats JSON NOT NULL,
    FOREIGN KEY (comment_id) REFERENCES comments(id)
);

-- Wstawianie danych do tabeli userinfo
INSERT INTO userinfo (email, password, role) VALUES
                                                 ('user1@example.com', 'password123', 1),
                                                 ('user2@example.com', 'securepass456', 2);

-- Wstawianie danych do tabeli userrole
INSERT INTO userrole (name) VALUES
                                ('admin'),
                                ('user');

-- Wstawianie danych do tabeli posts
INSERT INTO posts (text) VALUES
                             ('Dziś pogoda jest naprawdę piękna, słoneczko świeci, a niebo jest bezchmurne.'),
                             ('Ostatnio czytałem fascynującą książkę o historii Polski.'),
                             ('Planuję w najbliższym czasie zrobić sobie wycieczkę w góry.'),
                             ('Nie mogę się doczekać wakacji, aby odpocząć nad morzem.'),
                             ('Dzisiaj spotkałem starych przyjaciół na kawie, było miło porozmawiać.'),
                             ('Kocham oglądać zachody słońca, zawsze są takie magiczne.'),
                             ('Dziś zrobiłem pyszne ciasto czekoladowe, wszyscy w domu byli zachwyceni.'),
                             ('Właśnie skończyłem remontować salon, teraz jest tak przytulnie.'),
                             ('Planuję dzisiaj wieczorem pójść do kina na premierę nowego filmu.'),
                             ('Nie lubię poniedziałków, zawsze są takie ciężkie po weekendzie.');


-- Wstawianie danych do tabeli comments
INSERT INTO comments (text, id_post) VALUES
                                         ('Gre  at post!', 1),
                                         ('I agree!', 2);

-- Wstawianie danych do tabeli page
INSERT INTO page (post_list) VALUES
    ('{"ids": [1, 2]}'),
    ('{"ids": [10]}')
    ('{"ids": [3,5,6]}')
    ('{"ids": [7,8,9]}');

-- Wstawianie danych do tabeli postreactions
INSERT INTO postreactions (target_id, post_id, stats) VALUES
                                                          (1, 1, '{"likes": 10, "shares": 5}'),
                                                          (2, 2, '{"likes": 8, "shares": 3}');

-- Wstawianie danych do tabeli comreactions
INSERT INTO comreactions (target_id, comment_id, stats) VALUES
                                                            (1, 1, '{"likes": 5, "replies": 2}'),
                                                            (2, 2, '{"likes": 3, "replies": 1}');

