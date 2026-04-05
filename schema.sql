-- SQL скрипт для инициализации базы данных

CREATE DATABASE IF NOT EXISTS вашаБД; -- Замените 'вашаБД' на имя вашей базы данных;

USE вашаБД; -- Замените 'вашаБД' на имя вашей базы данных

CREATE TABLE IF NOT EXISTS articles (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    anons TEXT,
    full_text TEXT
);

