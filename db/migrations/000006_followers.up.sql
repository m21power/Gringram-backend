CREATE TABLE followers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    follower_ID INT NOT NULL,
    followee_ID INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (follower_ID, followee_ID),
    FOREIGN KEY (follower_ID) REFERENCES user(id),
    FOREIGN KEY (followee_ID) REFERENCES user(id)
);