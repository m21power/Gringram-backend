CREATE TABLE user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    bio TEXT,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    profile_ID INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (profile_ID) REFERENCES profile_image(id)
);