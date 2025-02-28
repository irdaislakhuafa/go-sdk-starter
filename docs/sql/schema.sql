CREATE DATABASE IF NOT EXISTS `starter_app`;
USE `starter_app`;

-- todo
CREATE TABLE `todo` (
 `id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
 `title` VARCHAR(255) NOT NULL,
 `description` TEXT NOT NULL,
 `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `created_by` VARCHAR(255) NOT NULL,
 `updated_at` TIMESTAMP NULL,
 `updated_by` VARCHAR(255) NULL,
 `deleted_at` TIMESTAMP NULL,
 `deleted_by` VARCHAR(255) NULL,
 `is_deleted` TINYINT NOT NULL DEFAULT 0
);
