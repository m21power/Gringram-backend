CREATE TABLE post_image (
    id INT AUTO_INCREMENT PRIMARY KEY,
    url VARCHAR(255) NOT NULL,
    post_ID INT NOT NULL,
    FOREIGN KEY (post_ID) REFERENCES post(id)
);