CREATE DATABASE IF NOT EXISTS `strider`;

USE `strider`;

CREATE TABLE IF NOT EXISTS `strider`.`posts` (
    `id` INT AUTO_INCREMENT NOT NULL,
    `uuid` VARCHAR(36) NOT NULL,
    `type` ENUM('original', 'repost', 'quote') NOT NULL,
    `text` VARCHAR(777) NOT NULL,
    `user_id` INT NOT NULL,
    `post_id` INT,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`post_id`) REFERENCES `posts`(`id`)
) ENGINE = INNODB;