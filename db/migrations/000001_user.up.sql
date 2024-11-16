CREATE TABLE `user` (
    id INT AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    bio VARCHAR(255) DEFAULT '',
    `password` VARCHAR(255) NOT NULL,
    profile_image_url VARCHAR(255) DEFAULT '',
    -- URL to the user's profile image
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);