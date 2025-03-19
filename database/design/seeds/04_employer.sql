/**
  Adding a dummy employer, to show the relationship
  between employers and employer-pension products

  Although it's not required for this project,
  we could have a user-employer table showing current and past employers for each user.
  (The user may have been with multiple employers who have Cushon pension products.)

  In any case, the relationship can be inferred through the user_account table
 */

INSERT INTO `cushondb`.`employer` (`title`)
    VALUES ('FreshCo Supermarkets Limited');
