-- trigger: sebelum data di tweets dihapus, maka hapus dulu data di retweets
DELIMITER $$
CREATE TRIGGER delete_all_data_tweet
BEFORE DELETE ON tweets FOR EACH ROW
BEGIN
-- declare variables
DECLARE v_tweet_id VARCHAR(255);
SET v_tweet_id=OLD.id;
-- trigger code
DELETE FROM retweets WHERE tweet_id=v_tweet_id;
END$$
DELIMITER ;
