package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var (
	DB *sql.DB
)

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "./db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}

	createUserTable := `
    CREATE TABLE IF NOT EXISTS user (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        email TEXT NOT NULL,
		password TEXT NOT NULL,
		role TEXT NOT NULL,
		status TEXT NOT NULL
    );
    `
	_, err = DB.Exec(createUserTable)
	if err != nil {
		log.Fatal(err)
	}

	createCategoryTable := `
    CREATE TABLE IF NOT EXISTS category (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
		description TEXT,
		img TEXT
    );
    `
	_, err = DB.Exec(createCategoryTable)
	if err != nil {
		log.Fatal(err)
	}

	createProductTable := `
    CREATE TABLE IF NOT EXISTS product (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		img TEXT,
		price INTEGER NOT NULL,
		stock INTEGER NOT NULL,
        category_id INTEGER,
    FOREIGN KEY(category_id) REFERENCES category(id)
    );
    `
	_, err = DB.Exec(createProductTable)
	if err != nil {
		log.Fatal(err)
	}
}
