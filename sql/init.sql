-- Adminer 4.6.3 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `account`;
CREATE TABLE `account` (
  `account_number` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `customer_id` int(10) unsigned NOT NULL,
  `opening_date` date NOT NULL,
  `closing_date` date DEFAULT NULL,
  `balance` decimal(17,2) NOT NULL DEFAULT '0.00',
  PRIMARY KEY (`account_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


DROP TABLE IF EXISTS `customer`;
CREATE TABLE `customer` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `first_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `last_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `join_date` date NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


DROP TABLE IF EXISTS `transaction`;
CREATE TABLE `transaction` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `account_number` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `description` varchar(500) COLLATE utf8_unicode_ci NOT NULL,
  `amount` decimal(17,2) NOT NULL,
  `dtype` enum('debit','credit') COLLATE utf8_unicode_ci NOT NULL,
  `date` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `account_number` (`account_number`),
  CONSTRAINT `transaction_ibfk_1` FOREIGN KEY (`account_number`) REFERENCES `account` (`account_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


DELIMITER ;;

CREATE TRIGGER `transaction_ai_balance` AFTER INSERT ON `transaction` FOR EACH ROW
IF NEW.dtype = "debit" THEN
  UPDATE account SET account.balance = account.balance + NEW.amount WHERE account.account_number = NEW.account_number;
ELSE
  UPDATE account SET account.balance = account.balance - NEW.amount WHERE account.account_number = NEW.account_number;
END IF;;

DELIMITER ;

-- 2018-09-06 17:24:37
