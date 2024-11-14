ALTER TABLE
    followers DROP FOREIGN KEY followers_ibfk_1;

ALTER TABLE
    followers DROP FOREIGN KEY followers_ibfk_2;

DROP TABLE IF EXISTS followers;