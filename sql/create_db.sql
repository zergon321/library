CREATE DATABASE IF NOT EXISTS library;

USE library;

CREATE TABLE IF NOT EXISTS `users` (
	`id` INT NOT NULL AUTO_INCREMENT,
    `nickname` VARCHAR(128) UNIQUE NOT NULL,
    `name` VARCHAR(128) NOT NULL,
    `surname` VARCHAR(128) NOT NULL,
    `patronim` VARCHAR(128) NOT NULL,
    `email` VARCHAR(128) UNIQUE NOT NULL,
    `password` VARCHAR(128) NOT NULL,
    PRIMARY KEY (`id`))
ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS `books` (
	`id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(512) NOT NULL,
    `author_name` VARCHAR(128) NOT NULL,
	`author_surname` VARCHAR(128) NOT NULL,
    `vendor_code` CHAR(32) UNIQUE NOT NULL,
    `price` REAL NOT NULL,
    PRIMARY KEY (`id`))
ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS `users_to_books` (
	`user_id` INT NOT NULL,
    `book_id` INT NOT NULL,
    `ordered` DATETIME NOT NULL,
    `delivered` DATETIME NULL,
    FOREIGN KEY (`user_id`) REFERENCES `library`.`users` (`id`) ON DELETE CASCADE,
    FOREIGN KEY (`book_id`) REFERENCES `library`.`books` (`id`) ON DELETE CASCADE)
ENGINE=InnoDB;