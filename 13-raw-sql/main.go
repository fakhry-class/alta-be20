package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
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

type TweetAccount struct {
	ID         string
	AccountId  uint
	Tweet      string
	SenderName string
}

func InitDB() (*sql.DB, error) {
	// DSN (Data Source Name):
	// <username>:<password>@tcp(<hostname>:<port-db>)/<db-name>

	var connectionString = os.Getenv("CONNECTION_DB")
	var db *sql.DB
	var err error
	// Get a database handle.
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Println("error open connection to db: ", err)
		return nil, err
	}

	// db.SetConnMaxLifetime(time.Minute * 3)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)

	// cek apakah sudah bisa connect ke db
	pingErr := db.Ping()
	if pingErr != nil {
		log.Println("err ping conenction: ", pingErr)
		return nil, pingErr
	}
	fmt.Println("success connect to db!")
	return db, nil
}

func main() {
	db, errInitDB := InitDB()
	if errInitDB != nil {
		log.Fatal("error connect to db ", errInitDB)
	}
	// var db *sql.DB
	//close connection
	defer db.Close()

	fmt.Println("Pilih menu:")
	fmt.Println("[1]: Read data")
	fmt.Println("[2]: Add data")
	fmt.Println("[3]: Delete data")
	fmt.Println("[0]: keluar aplikasi")
	var pilihan int
	fmt.Println("Masukkan angka sesuai pilihan menu:")
	fmt.Scanln(&pilihan)
	switch pilihan {
	case 1:
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

	case 2:
		// meminta inputan dari user
		fmt.Println("insert data")
		var newAccount = Account{}
		var join_date string
		fmt.Println("Masukkan username")
		fmt.Scanln(&newAccount.Username)
		fmt.Println("Masukkan name")
		fmt.Scanln(&newAccount.Name)
		fmt.Println("Masukkan bio")
		fmt.Scanln(&newAccount.Bio)
		fmt.Println("Masukkan tanggal join")
		fmt.Scanln(&join_date)
		fmt.Println("Masukkan alamat")
		fmt.Scanln(&newAccount.Location)
		fmt.Println("Masukkan email")
		fmt.Scanln(&newAccount.Email)
		fmt.Println("Masukkan password")
		fmt.Scanln(&newAccount.Password)

		//menjalakan query insert
		statement, errPrepare := db.Prepare("INSERT INTO accounts (username, name, bio, join_date, location, email, password) VALUES (?, ?,?,?,?,?,?)")
		if errPrepare != nil {
			log.Fatal("error prepare insert: ", errPrepare)
		}
		result, errInsert := statement.Exec(newAccount.Username, newAccount.Name, newAccount.Bio, join_date, newAccount.Location, newAccount.Email, newAccount.Password)
		if errInsert != nil {
			log.Fatal("error insert data", errInsert.Error())
		} else {
			// lastId, _ := result.LastInsertId()
			// fmt.Println("last id:", lastId)
			row, _ := result.RowsAffected()
			if row > 0 {
				fmt.Println("success insert data")
			} else {
				fmt.Println("failed insert data")
			}
		}

	case 3:
		fmt.Println("Masukkan email yang akan dihapus;")
		var emailInput string
		fmt.Scanln(&emailInput)
		//menjalakan query insert
		result, errDelete := db.Exec("DELETE FROM accounts WHERE email = ?", emailInput)
		if errDelete != nil {
			log.Fatal("error delete data", errDelete.Error())
		} else {
			row, _ := result.RowsAffected()
			if row > 0 {
				fmt.Println("success delete data")
			} else {
				fmt.Println("failed delete data")
			}
		}

	}

}
