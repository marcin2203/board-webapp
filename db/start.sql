-- Utwórz bazę danych db
CREATE DATABASE db;

-- Użyj bazy danych db
\c db

CREATE TABLE IF NOT EXISTS userrole (
    id SERIAL PRIMARY KEY,
    name VARCHAR(60) NOT NULL
    );

CREATE TABLE IF NOT EXISTS userinfo (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    nick VARCHAR(64),
    password VARCHAR(256) NOT NULL,
    role INT,
    FOREIGN KEY (role) REFERENCES userrole(id)
);

CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    text VARCHAR(1500) NOT NULL
);


CREATE TABLE IF NOT EXISTS tags (
    id SERIAL PRIMARY KEY,
    tag VARCHAR(64) NOT NULL
);

CREATE TABLE IF NOT EXISTS posts_tags (
    post_id INT,
    tag_id INT, 
    FOREIGN KEY (post_id) REFERENCES posts(id),
    FOREIGN KEY (tag_id) REFERENCES tags(id)
    );

CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    text VARCHAR(300) NOT NULL,
    post_id INT,
    FOREIGN KEY (post_id) REFERENCES posts(id)
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

-- Wstawianie danych do tabeli userrole
INSERT INTO userrole (name) VALUES
('admin'),
('user');

-- Wstawianie danych do tabeli userinfo
INSERT INTO userinfo (nick, email, password, role) VALUES
('adam1234', 'user1@example.com', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', 1),
('jola3333', 'user2@example.com', '03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4', 2);

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

INSERT INTO tags (tag) VALUES
('pogoda'),
('wakacje'),
('wolny-czas'),
('kot');

INSERT INTO posts_tags (post_id, tag_id) VALUES
(1,1), (1,3), (1,2), (3,2), (4,2), (4,3), (8,3), (9,3);

-- Wstawianie danych do tabeli comments
INSERT INTO comments (text, id_post) VALUES
('Gre  at post!', 1),
('I agree!', 2);

-- delete
-- Wstawianie danych do tabeli page
INSERT INTO page (post_list) VALUES
    ('{"ids": [7,8,9]}'),
    ('{"ids": [10]}'),
    ('{"ids": [3,5,6]}'),
    ('{"ids": [7,8,9]}');

-- Wstawianie danych do tabeli postreactions
INSERT INTO postreactions (target_id, post_id, stats) VALUES
(1, 1, '{"likes": 10, "shares": 5}'),
(2, 2, '{"likes": 8, "shares": 3}');

-- Wstawianie danych do tabeli comreactions
INSERT INTO comreactions (target_id, comment_id, stats) VALUES
(1, 1, '{"likes": 5, "replies": 2}'),
(2, 2, '{"likes": 3, "replies": 1}');

