use db_twitter;

-- DML
-- INSERT 

-- insert ke table accounts
INSERT INTO accounts(username, name, bio, join_date, location, email, password)
VALUES ("user1", "user 1", "lorem ipsum dolor sit amet", "2023-11-30", "Indonesia", "user1@mail.com", "qwerty");

INSERT INTO accounts(username, name, bio, join_date, location, email, password)
VALUES ("userA", "user A besar satu", "lorem ipsum dolor sit amet", "2023-11-30", "Indonesia", "usera@mail.com", "qwerty");

INSERT INTO accounts(username, name, bio, join_date, location, email, password)
VALUES ("userb", "user b besar", "lorem ipsum dolor sit amet", "2023-12-01", "Indonesia", "userb@mail.com", "qwerty");

INSERT INTO accounts(username, name, bio, join_date, location, email, password)
VALUES ("userc", "user c besar", "lorem ipsum dolor sit amet", "2023/12/01", "Indonesia", "userc@mail.com", "qwerty");

INSERT INTO accounts(username, name, bio, join_date, location, email, password)
VALUES ("userc1", "user c kecil", "lorem ipsum dolor sit amet", "2023/12/01", "Indonesia", "userc1@mail.com", "qwerty");

INSERT INTO accounts(username, name, join_date, location, email)
VALUES ("userd", "user d", "2023/12/01", "Indonesia", "userd@mail.com");


-- tambahan data setelah dihapus
INSERT INTO accounts(username, name, bio, join_date, location, email, password)
VALUES ("budi", "budi sudarsono", "lorem ipsum dolor sit amet", "2023/12/05", "Indonesia", "budi@mail.com", "qwerty");
INSERT INTO accounts(username, name, bio, join_date, location, email, password)
VALUES ("andi", "andi f noya", "lorem ipsum dolor sit amet", "2023/12/06", "surabaya", "andi@mail.com", "qwerty");
INSERT INTO accounts(username, name, bio, join_date, location, email, password)
VALUES ("rudi", "rudi rahadian", "lorem ipsum dolor sit amet", "2023/12/06", "surabaya", "rudi@mail.com", "qwerty");
INSERT INTO accounts(username, name, bio, join_date, location, email, password)
VALUES ("reza", "reza rahadian", "lorem ipsum dolor sit amet", "2023/12/06", "jogjakarta", "reza@mail.com", "qwerty");


-- insert ke table tweets 
INSERT INTO tweets (id, account_id, tweet, comment_count, retweet_count, like_count)
VALUES ("X0004", 10, "halo ini tweet pertama", 0, 0, 0);
INSERT INTO tweets (id, account_id, tweet, comment_count, retweet_count, like_count)
VALUES ("X0002", 4, "halo ini tweet kedua", 0, 0, 0);

INSERT INTO tweets (id, account_id, tweet, comment_count, retweet_count, like_count)
VALUES ("X0003", 100, "halo ini tweet seratus", 0, 0, 0);

INSERT INTO tweets (id, account_id, tweet, comment_count, retweet_count, like_count)
VALUES ("abc0003", 2, "halo ini tweet dua satu", 0, 0, 0);

INSERT INTO tweets (id, account_id, tweet, comment_count, retweet_count, like_count)
VALUES ("abc0005", 3, "halo ini tweet tiga satu", 0, 0, 0);



-- SELECT 
SELECT * FROM accounts;

SELECT * FROM tweets;

-- membaca seluruh data di spesifik kolom dari satu table
SELECT id, username, email from accounts;

-- membaca spesifik data tertentu dari satu table dengan where clause
SELECT id, username,name, email from accounts WHERE id = 1;
SELECT id, username, email from accounts WHERE username = "userb";
SELECT * from accounts WHERE username = "userb";
SELECT * from accounts WHERE join_date  = "2023-12-01";
SELECT * from accounts WHERE password IS NOT NULL ;
SELECT * from accounts WHERE password IS NULL ;


-- UPDATE 
-- melakukan edit/update data

UPDATE accounts SET 
name = "user 1 edit"
WHERE id = 1;

UPDATE accounts SET 
bio = "ini bio edit",
location = "jakarta"
WHERE id = 2;

-- DELETE 
-- menghapus data

DELETE FROM accounts WHERE id = 7;
DELETE FROM accounts WHERE id = 2;

-- FK
-- RESTRICT = kita tidak bisa menghapus data parent ketika masih dibutuhkan di table child.
-- cara hapus: hapus dulu child nya, baru bisa hapus parent.

-- CASCADE = ketika kita hapus data parent, maka yang child akan ikut kehapus

DELETE FROM tweets WHERE id ="X0003";
SELECT * from tweets t ;
SELECT * FROM accounts a ;
DELETE FROM accounts WHERE id = 2;

DELETE FROM accounts WHERE id = 1;


-- LIKE /  BETWEEN
SELECT * from accounts;
SELECT * from accounts  WHERE name LIKE "%besar";
SELECT * from accounts  WHERE name LIKE "%besar%";

SELECT * FROM accounts WHERE id BETWEEN 5 and 10;
SELECT * from accounts WHERE join_date BETWEEN "2023-12-01" and "2023-12-10";

-- AND OR 
SELECT * from accounts;
SELECT * from accounts WHERE username LIKE "%di%";
SELECT * from accounts WHERE username LIKE "%di%" AND location = 'SURABAYA';
SELECT * from accounts WHERE username LIKE "%di%" OR location = 'jogjakarta' OR id BETWEEN 5 and 10 ;

-- ORDER BY
-- ASC: kecil ke besar
-- DESC : besar ke kecil
SELECT * from accounts;
SELECT * FROM accounts order by id DESC ;
SELECT * FROM accounts order by join_date ASC ;
SELECT * from accounts WHERE username LIKE "%di%" order by id DESC;

-- LIMIT
SELECT * from accounts LIMIT 2;
SELECT * from accounts WHERE username LIKE "%di%" order by id DESC LIMIT 1;
SELECT * from accounts WHERE username LIKE "%di%" order by id ASC LIMIT 1;


-- JOIN

select * FROM tweets t ;
SELECT * from accounts;

-- insert account id nya null
INSERT INTO tweets (id, tweet, comment_count, retweet_count, like_count)
VALUES ("abc0004", "halo ini kosongan", 0, 0, 0);

-- INNER JOIN
SELECT tweets.id, tweets.account_id , tweets.tweet, accounts.name , accounts.username, tweets.created_at FROM tweets
INNER JOIN accounts ON tweets.account_id = accounts.id ;

-- memberikan alias ke tabel dan kolom
SELECT t.id, t.account_id , t.tweet, a.name as sender_name , a.username, t.created_at FROM tweets t
INNER JOIN accounts a ON t.account_id = a.id ;

SELECT tweets.id, tweets.account_id , tweets.tweet, tweets.comment_count, accounts.name , accounts.username, tweets.created_at FROM tweets
INNER JOIN accounts ON tweets.account_id = accounts.id 
WHERE tweets.comment_count < 20;

-- LEFT JOIN
SELECT tweets.id, tweets.account_id , tweets.tweet, accounts.name , accounts.username, tweets.created_at FROM tweets
LEFT JOIN accounts ON tweets.account_id = accounts.id ;

SELECT tweets.id, tweets.account_id , tweets.tweet, accounts.name , accounts.username, tweets.created_at FROM accounts
LEFT JOIN tweets ON tweets.account_id = accounts.id ;

-- RIGHT JOIN
SELECT tweets.id, tweets.account_id , tweets.tweet, accounts.name , accounts.username, tweets.created_at FROM tweets
RIGHT JOIN accounts ON tweets.account_id = accounts.id ;


SELECT * FROM retweets r ;
select * FROM tweets t ;
SELECT * from accounts;

INSERT INTO retweets (account_id, tweet_id)
VALUES (4,"abc0003"),
(8,"X0001"),
(9,"X0001"),
(8,"X0002");

-- JOIN lebih dari 2 table

-- menampilkan data tweet, beserta data nama yang nge tweet
SELECT 
tweets.id , tweets.tweet, tweets.account_id as sender_id, 
accounts.username , accounts.name
from tweets
INNER JOIN accounts ON tweets.account_id = accounts.id;

-- menampilkan data tweet, beserta data nama yang nge tweet, dan account_id yang nge retweet
SELECT 
tweets.id , tweets.tweet, tweets.account_id as sender_id, 
accounts.username , accounts.name, 
retweets.account_id as retweet_account_id
from tweets
INNER JOIN accounts ON tweets.account_id = accounts.id 
INNER JOIN retweets ON tweets.id = retweets.tweet_id;

-- menampilkan data tweet, beserta data nama yang nge tweet, dan account_id dan nama yang nge retweet
SELECT 
tweets.id , tweets.tweet, tweets.account_id as sender_id, 
accounts.username , accounts.name, 
retweets.account_id as retweet_account_id, ar.name as retweet_name
from tweets
INNER JOIN accounts ON tweets.account_id = accounts.id 
INNER JOIN retweets ON tweets.id = retweets.tweet_id 
INNER JOIN accounts ar ON retweets.account_id = ar.id ;


-- UNION
select id from accounts 
UNION
select id from tweets ;

-- distinct
SELECT * from accounts ;
SELECT DISTINCT location from accounts a ;

-- AGGREGATION
-- MIN
SELECT * from accounts ;
SELECT MIN(id) as min_id from accounts a ; 
SELECT MIN(join_date) from accounts a ; 

SELECT * FROM tweets t ;
SELECT MIN(comment_count) from tweets t ; 

-- MAX
SELECT * from accounts ;
SELECT MAX(id) as max_id from accounts a ; 
SELECT MAX(join_date) from accounts a ; 

SELECT * FROM tweets t ;
SELECT MAX(comment_count) from tweets t ; 

-- SUM
SELECT * FROM tweets t ;
SELECT SUM(comment_count) from tweets t ;
SELECT SUM(comment_count) from tweets t WHERE account_id = 4;

-- menampilkan jumlah keseluruhan comment yang didapat oleh masing" akun. 
SELECT account_id, SUM(comment_count) 
from tweets t
group by account_id ;

SELECT t.account_id, a.name, SUM(t.comment_count) 
from tweets t
INNER JOIN accounts a ON t.account_id = a.id 
group by t.account_id ;

SELECT t.account_id, a.name, SUM(t.comment_count) as sum_comment, SUM(t.like_count) as sum_like
from tweets t
INNER JOIN accounts a ON t.account_id = a.id 
group by t.account_id ;

-- AVG / Average / rata-rata
SELECT AVG(comment_count) from tweets t ;
SELECT AVG(comment_count) from tweets t WHERE account_id = 4;

SELECT t.account_id, a.name,SUM(t.comment_count) , AVG(t.comment_count) 
from tweets t
INNER JOIN accounts a ON t.account_id = a.id 
group by t.account_id ;

-- COUNT
SELECT  * FROM tweets t2 ;
SELECT COUNT(id) from tweets t ; 
SELECT COUNT(id) from tweets t WHERE account_id = 4; 

-- menampilkan banyaknya tweet yang telah dibuat oleh masing" akun
SELECT t.account_id, a.name, COUNT(t.id) , SUM(t.comment_count) , AVG(t.comment_count) 
from tweets t
INNER JOIN accounts a ON t.account_id = a.id 
group by t.account_id ;

-- HAVING
-- menampilkan data akun yang sudah nge tweet lebih dari 1
SELECT t.account_id, a.name, COUNT(t.id) , SUM(t.comment_count) , AVG(t.comment_count) 
from tweets t
INNER JOIN accounts a ON t.account_id = a.id 
group by t.account_id 
HAVING COUNT(t.id) > 1;

-- menampilkan data akun yang jumlah comment nya >= 20
SELECT t.account_id, a.name, COUNT(t.id) , SUM(t.comment_count) , AVG(t.comment_count) 
from tweets t
INNER JOIN accounts a ON t.account_id = a.id 
group by t.account_id 
HAVING SUM(t.comment_count) >= 20
ORDER BY SUM(t.comment_count) DESC;

SELECT t.account_id, a.name, COUNT(t.id) , SUM(t.comment_count) , AVG(t.comment_count) 
from tweets t
INNER JOIN accounts a ON t.account_id = a.id 
group by t.account_id 
HAVING SUM(t.comment_count) >= 20
ORDER BY SUM(t.comment_count) DESC
LIMIT 1;


-- SUBQUERY
-- a. mendapatkan list akun id yang sudah/pernah ngetweet
SELECT account_id from tweets t WHERE account_id is not null;
SELECT DISTINCT account_id from tweets t WHERE account_id is not null;

-- b. menampilkan data akunc
SELECT  * from accounts a ;

-- menampilkan data akun yang pernah ngetweet
SELECT * FROM accounts a WHERE id IN (SELECT DISTINCT account_id from tweets t WHERE account_id is not null);
-- SELECT * FROM accounts a WHERE id IN (3,4,10);

-- menampilkan data akun yang belum pernah ngetweet
SELECT * FROM accounts a WHERE id NOT IN (SELECT DISTINCT account_id from tweets t WHERE account_id is not null);

-- menampilkan data akun yang pernah ngetweet menggunakan pendekatan join
SELECT DISTINCT a.id ,a.username ,a.name 
FROM accounts a
inner join tweets t ON a.id = t.account_id ;


-- FUNCTION
-- fungsi untuk menghitung banyaknya tweet dari salah satu akun
-- NB: buat file baru untuk execute/ membuat function di dbeaver
DELIMITER $$
CREATE FUNCTION sf_hitung_tweet (account_id_param int) RETURNS INT DETERMINISTIC
BEGIN
DECLARE total INT;
SELECT COUNT(id) INTO total FROM tweets WHERE account_id = account_id_param;
RETURN total;
END$$
DELIMITER ;

-- melihat function yang ada di sebuah db
SHOW FUNCTION STATUS WHERE db = 'db_twitter';

-- menjalankan function
SELECT sf_hitung_tweet(4);

-- menghapus function
DROP function `db_twitter`.`sf_hitung_tweet`;


-- DEMO TRIGGER
-- NB: buat dulu func triggernya, jika sudah berhasil, maka jalankan perintah delete tweets

SELECT * from retweets r ;
DELETE from tweets WHERE id ="abc0003";

-- menghapus trigger
DROP TRIGGER db_twitter.delete_all_data_tweet;



SELECT id, username, name, bio, join_date, location, email, password, created_at FROM accounts;













