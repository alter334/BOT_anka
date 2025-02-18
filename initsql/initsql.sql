CREATE DATABASE IF NOT EXISTS `anka`;
USE `anka`;
DROP USER IF EXISTS 'anka'@'%';
CREATE USER 'anka'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON anka.* TO 'anka'@'%';
FLUSH PRIVILEGES;
-- 以下traQer関連

CREATE TABLE IF NOT EXISTS `ankaChannel` (
  `Id` char(36) NOT NULL,
  `channelMessageCount` int(11) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `ankaDB` (
  `Id` char(36) NOT NULL,
  `channelId` char(36) NOT NULL,
  `ankaInvokeMessageNum` int(11) NOT NULL,
  PRIMARY KEY (`Id`),
  FOREIGN KEY (`channelId`) REFERENCES `ankaChannel`(`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


