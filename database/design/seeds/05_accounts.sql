
/**
  User 1 has a standard employer pension with a balance
  That user's pension is invested in a single fund, as per the product's default
  But this design allows for different users, when policies change, to have different fund combinations
  even in their workplace pension
*/
INSERT INTO `cushondb`.`user_account` (`user_id`, `product_id`, `total_investment`, `current_balance`)
    VALUES ('1', '1', '6000', '6795.42');
INSERT INTO `cushondb`.`user_account_funds` (`fund_id`, `user_account_id`, `weight_pc`)
    VALUES ('1', '1', '100');
/* Simulate a monthly pension contribution totaling 6k */
INSERT INTO `cushondb`.`user_account_transaction` (`account_id`, `created`, `amount`)
VALUES
    ('1', '2023-03-01', '250'),
    ('1', '2023-04-01', '250'),
    ('1', '2023-05-01', '250'),
    ('1', '2023-06-01', '250'),
    ('1', '2023-07-01', '250'),
    ('1', '2023-08-01', '250'),
    ('1', '2023-09-01', '250'),
    ('1', '2023-10-01', '250'),
    ('1', '2023-11-01', '250'),
    ('1', '2023-12-01', '250'),
    ('1', '2024-01-01', '250'),
    ('1', '2024-02-01', '250'),
    ('1', '2024-03-01', '250'),
    ('1', '2024-04-01', '250'),
    ('1', '2024-05-01', '250'),
    ('1', '2024-06-01', '250'),
    ('1', '2024-07-01', '250'),
    ('1', '2024-08-01', '250'),
    ('1', '2024-09-01', '250'),
    ('1', '2024-10-01', '250'),
    ('1', '2024-11-01', '250'),
    ('1', '2024-12-01', '250'),
    ('1', '2025-01-01', '250'),
    ('1', '2025-02-01', '250');


/** User 2 has no product relationship (yet), except a platform login account */

