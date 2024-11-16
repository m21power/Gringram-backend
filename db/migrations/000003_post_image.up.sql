CREATE TABLE `post_image` (
    id INT AUTO_INCREMENT PRIMARY KEY,
    post_id INT NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    -- URL to the image
    FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE
);