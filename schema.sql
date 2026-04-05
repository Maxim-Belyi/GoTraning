CREATE DATABASE IF NOT EXISTS golangBd; 

USE golangBd; 

CREATE TABLE IF NOT EXISTS articles (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    anons TEXT,
    full_text TEXT
);

