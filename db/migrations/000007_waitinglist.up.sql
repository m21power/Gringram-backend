CREATE TABLE waitinglist (
    id INT AUTO_INCREMENT PRIMARY KEY,
    post_id INT NOT NULL,
    `status` ENUM('pending', 'approved', 'rejected') DEFAULT 'pending',
    FOREIGN KEY (post_id) REFERENCES posts(id)
);