
CREATE DATABASE jwtLogin;

USE `jwtLogin`;

CREATE TABLE IF NOT EXISTS `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` text,
  `email` text,
  `userName` text,
  `refToken` text,
  `isActive` boolean NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;