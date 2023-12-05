DELIMITER $$
CREATE FUNCTION sf_hitung_tweet2 (account_id_param int) RETURNS INT DETERMINISTIC
BEGIN
DECLARE total INT;
SELECT COUNT(id) INTO total FROM tweets WHERE account_id = account_id_param;
RETURN total;
END$$
DELIMITER ;