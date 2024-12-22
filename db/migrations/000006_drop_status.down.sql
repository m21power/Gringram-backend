ALTER TABLE posts ADD COLUMN `status` ENUM('pending', 'approved', 'rejected') DEFAULT 'pending';
