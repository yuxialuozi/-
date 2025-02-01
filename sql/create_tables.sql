-- Create video table
CREATE TABLE `video`
(
    `id`         INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` DATETIME DEFAULT NULL,
    `user_id`    INT UNSIGNED NOT NULL,
    `title`      VARCHAR(255) NOT NULL,
    `file_name`  VARCHAR(255) NOT NULL,
    `cover_name` VARCHAR(255) NOT NULL
);

-- Create user table
CREATE TABLE `user`
(
    `id`               INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `created_at`       DATETIME     DEFAULT CURRENT_TIMESTAMP,
    `updated_at`       DATETIME     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`       DATETIME     DEFAULT NULL,
    `username`         VARCHAR(255) NOT NULL,
    `password`         VARCHAR(255) NOT NULL,
    `avatar`           VARCHAR(255) DEFAULT NULL,
    `background_image` VARCHAR(255) DEFAULT NULL,
    `signature`        VARCHAR(255) DEFAULT NULL
);

-- Create userToken table
CREATE TABLE `user_token`
(
    `id`       INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `token`    VARCHAR(255) NOT NULL,
    `username` VARCHAR(255) NOT NULL,
    `user_id`  INT UNSIGNED NOT NULL,
    `role`     VARCHAR(100) NOT NULL
);

-- Create relation table
CREATE TABLE `relation`
(
    `id`         INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` DATETIME DEFAULT NULL,
    `user_id`    INT UNSIGNED NOT NULL,
    `target_id`  INT UNSIGNED NOT NULL
);

-- Create favorite table
CREATE TABLE `favorite`
(
    `id`         INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` DATETIME DEFAULT NULL,
    `user_id`    INT UNSIGNED NOT NULL,
    `video_id`   INT UNSIGNED NOT NULL
);

-- Add indexes for performance (without foreign keys)
CREATE INDEX `idx_video_user_id` ON `video` (`user_id`);
CREATE INDEX `idx_user_token_user_id` ON `user_token` (`user_id`);
CREATE INDEX `idx_relation_user_id` ON `relation` (`user_id`);
CREATE INDEX `idx_relation_target_id` ON `relation` (`target_id`);
CREATE INDEX `idx_favorite_user_id` ON `favorite` (`user_id`);
CREATE INDEX `idx_favorite_video_id` ON `favorite` (`video_id`);
