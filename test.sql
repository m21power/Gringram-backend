CREATE TABLE `user` (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    bio VARCHAR(255) NOT NULL DEFAULT '',
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS `user`;

CREATE TABLE profile_image (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    image_data BLOB NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(id)
);

-- droping foreign key forget
DROP TABLE IF EXISTS profile_image;

CREATE TABLE post (
    id INT AUTO_INCREMENT PRIMARY KEY,
    content TEXT NOT NULL,
    user_ID INT NOT NULL,
    like_count INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_ID) REFERENCES user(id)
);

ALTER TABLE
    post DROP FOREIGN KEY post_ibfk_1;

DROP TABLE IF EXISTS post;

CREATE TABLE comment (
    id INT AUTO_INCREMENT PRIMARY KEY,
    content TEXT NOT NULL,
    user_ID INT NOT NULL,
    post_ID INT NOT NULL,
    like_count INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_ID) REFERENCES user(id),
    FOREIGN KEY (post_ID) REFERENCES post(id)
);

ALTER TABLE
    comment DROP FOREIGN KEY comment_ibfk_1;

ALTER TABLE
    comment DROP FOREIGN KEY comment_ibfk_2;

DROP TABLE IF EXISTS comment;

CREATE TABLE followers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    follower_ID INT NOT NULL,
    followee_ID INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (follower_ID, followee_ID),
    FOREIGN KEY (follower_ID) REFERENCES user(id),
    FOREIGN KEY (followee_ID) REFERENCES user(id)
);

ALTER TABLE
    followers DROP FOREIGN KEY followers_ibfk_1;

ALTER TABLE
    followers DROP FOREIGN KEY followers_ibfk_2;

DROP TABLE IF EXISTS followers;