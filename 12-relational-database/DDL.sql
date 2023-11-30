-- DDL
-- membuat database
create database db_twitter;
-- menampilkan semua db yang ada di mysql server
show databases;

-- memilih db yang akan digunakan / di manage
USE db_twitter;

-- membuat table accounts
CREATE TABLE accounts(
id int primary key auto_increment,
username varchar(100) not null unique,
name varchar(255),
bio text,
join_date date,
location varchar(255),
email varchar(255) not null unique,
password varchar(255),
created_at datetime DEFAULT CURRENT_TIMESTAMP
);

-- membuat table tweets
CREATE TABLE tweets(
id varchar(255) primary key,
account_id int,
tweet varchar(140),
comment_count int,
retweet_count int,
like_count int,
created_at datetime default current_timestamp,
updated_at datetime default current_timestamp,
deleted_at datetime,
constraint fk_tweet_account FOREIGN KEY (account_id) REFERENCES accounts(id)
);

CREATE TABLE retweets(
id int primary key auto_increment,
account_id int,
tweet_id varchar(255),
created_at datetime default current_timestamp,
updated_at datetime default current_timestamp,
deleted_at datetime,
constraint fk_retweet_account FOREIGN KEY (account_id) REFERENCES accounts(id),
constraint fk_retweet_tweet FOREIGN KEY(tweet_id) REFERENCES tweets(id)
);

-- menampilkan seluruh table
show tables;
-- menampilkan semua kolom di spesifik table
show columns from dummys;

create table dummys(
id int primary key,
name varchar(255)
);

-- memodifikasi table
-- menambahkan kolom
ALTER table dummys 
ADD COLUMN description varchar(255);

-- mengubah kolom
ALTER TABLE dummys 
MODIFY COLUMN description text;

-- menghapus kolom
ALTER TABLE dummys 
DROP COLUMN description;

-- alter foreign key
ALTER TABLE dummys
ADD CONSTRAINT fk_dummys_account
FOREIGN KEY (account_id) REFERENCES accounts(id);

ALTER TABLE dummys
DROP FOREIGN KEY fk_dummys_account;

-- rename nama table
rename table dummys to dummys_detail;

-- menghapus table
drop table dummys_detail;

-- menghapus database
drop database db_twitter;







