# source + .sql 文件路径
use debate;
CREATE TABLE user (
    uid INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    rid INT UNSIGNED NOT NULL DEFAULT 1,
    username VARCHAR(50) NOT NULL DEFAULT 'Visitor',
    password VARCHAR(50) NOT NULL,
    phone VARCHAR(30) NOT NULL,
    email VARCHAR(30) NOT NULL,
    img  VARCHAR(33) NOT NULL,
    role INT NOT NULL DEFAULT 2,
    score INT UNSIGNED NOT NULL DEFAULT 0
);
INSERT INTO user VALUES (1,1,'peterliang','ULwH1CwT3mCPMQ==','123456789','12345@qq.com','',1,0);

# 记录辩论的双方
-- CREATE TABLE debate (
--     id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
--     yid INT UNSIGNED NOT NULL,
--     nid INT UNSIGNED NOT NULL
-- );

CREATE TABLE debate (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(33) NOT NULL,
    positive_username VARCHAR(33),
    negative_username VARCHAR(33)
)