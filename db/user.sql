CREATE USER 'yofio_user'@'localhost' IDENTIFIED BY '2021Yofio.5347';
GRANT ALL PRIVILEGES ON yofio.* TO 'yofio_user'@'localhost';
flush privileges;
