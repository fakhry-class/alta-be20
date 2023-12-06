package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Account struct {
	ID        uint
	Username  string
	Name      string
	Bio       string
	JoinDate  time.Time
	Location  string
	Email     string
	Password  string
	CreatedAt time.Time
}

func main() {
	// DSN (Data Source Name):
	// <username>:<password>@tcp(<hostname>:<port-db>)/<db-name>
	var connectionString = "root:qwerty123@tcp(localhost:3306)/db_twitter?parseTime=true"
	var db *sql.DB
	var err error
	// Get a database handle.
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("error open connection to db: ", err)
	}

	// db.SetConnMaxLifetime(time.Minute * 3)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("err ping conenction: ", pingErr)
	}
	fmt.Println("success connect to db!")

	//close connection
	defer db.Close()

	// READ / SELECT data
	// var untuk menyimpan data yang dibaca di query select
	var accounts []Account

	// menjalankan perintah query select
	rows, errSelect := db.Query("SELECT id, username, name, bio, join_date, location, email, password, created_at FROM accounts")
	// handle ketika terjadi error saat menjalankan select
	if errSelect != nil {
		log.Fatal("error run query select ", errSelect.Error())
	}

	//proses membaca per baris
	// menggunakan for karena data yang akan dibaca itu banyak
	for rows.Next() {
		var dataAccount Account // var untuk menyimpan data akun per baris
		// var joinDate string
		// proses scan dan mapping data ke var dataAccount
		errScan := rows.Scan(&dataAccount.ID, &dataAccount.Username, &dataAccount.Name, &dataAccount.Bio, &dataAccount.JoinDate, &dataAccount.Location, &dataAccount.Email, &dataAccount.Password, &dataAccount.CreatedAt)
		if errScan != nil {
			log.Fatal("error scan select", errScan.Error())
		}
		// fmt.Println("joinDate:", joinDate)
		// memasukkan dataAccount ke accounts
		accounts = append(accounts, dataAccount)
	}

	for _, v := range accounts {
		fmt.Printf("username: %v, join-date: %v\n", v.Username, v.JoinDate)
	}
}
