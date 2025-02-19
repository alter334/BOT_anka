CREATE DATABASE IF NOT EXISTS `ankaDB`;
USE `ankaDB`;
DROP USER IF EXISTS 'ankaDB'@'%';
CREATE USER 'ankaDB'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON ankaDB.* TO 'ankaDB'@'%';
FLUSH PRIVILEGES;
-- 以下traQer関連
CREATE TABLE IF NOT EXISTS `anka` (
  `Id` char(36) NOT NULL,
  `originMessageId` char(36) NOT NULL,
  `channelId` char(36) NOT NULL,
  `ankaInvokeMessageCount` int(11) NOT NULL,
  Primary key (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


