-- Use InnoDB and utf8mb4 charset for all tables
CREATE TABLE IF NOT EXISTS `user` (
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `name_title` VARCHAR(16) NOT NULL,
    `name_first` VARCHAR(255) NOT NULL,
    `name_last` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL UNIQUE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
    COMMENT = 'For simplicity we assume "user" is also a "Person or individual"';

CREATE TABLE IF NOT EXISTS `employer` (
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `uuid` CHAR(36) NOT NULL UNIQUE DEFAULT (UUID()),
    `title` VARCHAR(255) NOT NULL UNIQUE COMMENT 'Employer name'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
    COMMENT = 'Simple example table for employer, AKA company or B2B entity customers';

CREATE TABLE IF NOT EXISTS `product` (
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `uuid` CHAR(36) NOT NULL UNIQUE DEFAULT (UUID()),
    `title` VARCHAR(255) NOT NULL COMMENT 'Product name',
    `type` ENUM('direct','employer') NOT NULL
        COMMENT 'Type of product. Direct is a retail product, employer is a B2B product',
    `description` TEXT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
    COMMENT = 'Product lists the available products or account types for direct and employer customers'
;

CREATE TABLE IF NOT EXISTS `user_account` (
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `user_id` INT UNSIGNED NOT NULL,
    `product_id` INT UNSIGNED NOT NULL,
    `total_investment` DECIMAL(10,2) NOT NULL DEFAULT 0 CHECK (`total_investment` >= 0),
    `current_balance` DECIMAL(10,2) NOT NULL DEFAULT 0 CHECK (`current_balance` >= 0)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `user_account_funds` (
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `fund_id` INT UNSIGNED NOT NULL,
    `user_account_id` INT UNSIGNED NOT NULL,
    `weight_pc` DECIMAL(10,2) NOT NULL DEFAULT 100 CHECK (`weight_pc` BETWEEN 0 AND 100)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `user_account_transaction` (
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `account_id` INT UNSIGNED NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `amount` DECIMAL(10,2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `fund` (
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `uuid` CHAR(36) NOT NULL UNIQUE DEFAULT (UUID()),
    `title` VARCHAR(255) NOT NULL UNIQUE,
    `description` TEXT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `employer_product` (
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `employer_id` INT UNSIGNED NOT NULL,
    `product_id` INT UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Indexes
CREATE INDEX `idx_user_email` ON `user` (`email`);

CREATE INDEX `idx_employer_uuid` ON `employer` (`uuid`);
CREATE INDEX `idx_employer_title` ON `employer` (`title`);

CREATE INDEX `idx_product_uuid` ON `product` (`uuid`);
CREATE UNIQUE INDEX `idx_product_title_type` ON `product` (`title`, `type`);

CREATE UNIQUE INDEX `idx_user_product` ON `user_account` (`user_id`, `product_id`);
CREATE INDEX `idx_user_account_user` ON `user_account` (`user_id`);

CREATE UNIQUE INDEX `idx_unique_user_fund` ON `user_account_funds` (`user_account_id`, `fund_id`);

CREATE UNIQUE INDEX `idx_unique_transaction_time` ON `user_account_transaction` (`account_id`, `amount`, `created`);

CREATE INDEX `idx_fund_uuid` ON `fund` (`uuid`);
CREATE INDEX `idx_fund_title` ON `fund` (`title`);

CREATE INDEX `idx_employer_product` ON `employer_product` (`employer_id`);
CREATE UNIQUE INDEX `idx_employer_product_unique` ON `employer_product` (`employer_id`, `product_id`);

-- Foreign Keys
ALTER TABLE `user_account`
    ADD CONSTRAINT `fk_user_account_user`
        FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
            ON DELETE CASCADE;

ALTER TABLE `user_account`
    ADD CONSTRAINT `fk_user_account_product`
        FOREIGN KEY (`product_id`) REFERENCES `product` (`id`)
            ON DELETE CASCADE;

ALTER TABLE `user_account_funds`
    ADD CONSTRAINT `fk_user_account_funds_fund`
        FOREIGN KEY (`fund_id`) REFERENCES `fund` (`id`)
            ON DELETE CASCADE;

ALTER TABLE `user_account_funds`
    ADD CONSTRAINT `fk_user_account_funds_user_account`
        FOREIGN KEY (`user_account_id`) REFERENCES `user_account` (`id`)
            ON DELETE CASCADE;

ALTER TABLE `user_account_transaction`
    ADD CONSTRAINT `fk_user_account_transaction_account`
        FOREIGN KEY (`account_id`) REFERENCES `user_account` (`id`)
            ON DELETE CASCADE;

ALTER TABLE `employer_product`
    ADD CONSTRAINT `fk_employer_product_employer`
        FOREIGN KEY (`employer_id`) REFERENCES `employer` (`id`)
            ON DELETE CASCADE;

ALTER TABLE `employer_product`
    ADD CONSTRAINT `fk_employer_product_product`
        FOREIGN KEY (`product_id`) REFERENCES `product` (`id`)
            ON DELETE CASCADE;
