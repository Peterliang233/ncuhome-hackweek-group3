# use debate;
# source + .sql 文件路径
CREATE TABLE user (
    uid INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    rid INT UNSIGNED NOT NULL DEFAULT 1,
    username VARCHAR(50) DEFAULT 'null',
    password VARCHAR(50) NOT NULL,
    phone VARCHAR(30),
    email VARCHAR(30) NOT NULL,
    img  VARCHAR(33),
    role INT NOT NULL DEFAULT 2,
    score INT UNSIGNED NOT NULL DEFAULT 0
);
CREATE TABLE identity (
    id INT UNSIGNED NOT NULL PRIMARY KEY,
    name VARCHAR(33) NOT NULL
);
INSERT INTO user VALUES (1,1,'peterliang','123456abc','123456789','12345@qq.com',1,0);
INSERT INTO identity VALUES (1,'学录');
INSERT INTO identity VALUES (2,'学士');
INSERT INTO identity VALUES (3,'大学士');