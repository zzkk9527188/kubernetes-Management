create database test if not exists;
use test;
CREATE TABLE `users` (
                         `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                         `username` varchar(255) NOT NULL,
                         `password` varchar(255) NOT NULL,
                         `created_at` datetime(3) NULL,
                         `updated_at` datetime(3) NULL,
                         `deleted_at` datetime(3) NULL,
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `username` (`username`),
                         INDEX `idx_users_deleted_at` (`deleted_at`)
);