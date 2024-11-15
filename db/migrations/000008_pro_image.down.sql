-- droping foreign key forget
ALTER TABLE
    profile_image DROP FOREIGN KEY profile_image_ibfk_1;

DROP TABLE IF EXISTS profile_image;