CREATE DATABASE IF NOT EXISTS `anka`;
USE `anka`;
DROP USER IF EXISTS 'anka'@'%';
CREATE USER 'anka'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON anka.* TO 'anka'@'%';
FLUSH PRIVILEGES;
-- 以下traQer関連
CREATE TABLE IF NOT EXISTS `ankaDB` (
  `ankaId` char(36) NOT NULL,
  `ankaInvokeMessageNum` int(11) NOT NULL,
  PRIMARY KEY (`ankaId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
