package mysqldatabase

import (
	"database/sql"
	"log"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func MySqlDatabase() {
	// Configure the database connection (MUST check for errors)
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/root?parseTime=true") // root here represents, username, password & dbname

	if err != nil {
		log.Fatal(err) // log the error
	}

	// initialize first connection to the database (Make sure to check the error), see if everything works correctly
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// create a new table
	{
		query :=
			`CREATE TABLE  users (
				id INT AUTO_INCREMENT,
				username TEXT NOT NULL,
				password TEXT NOT NULL,
				created_at DATETIME,
				PRIMARY KEY (id)
			)`

			// Execute the query in the DB, check err to ensure there was none
			if _, err := db.Exec(query); err != nil {
				log.Fatal(err)
			}
	}

	// Insert a new user
	{
		username := "johndoe"
		password := "secret"
		createdAt := time.Now()

		// insert data into the table, returns result & a possible error
		// result contains information about the last inserted id(auto-generated)
		// & a count of the rows affected by the query
		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
		if err != nil {
				log.Fatal(err)
		}

		// grabs the newly created id for your user
		userId, err := result.LastInsertId()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(userId)
}

// Query a single user
{
	// Create variables to store the data
		var (
				id        int
				username  string
				password  string
				createdAt time.Time
		)

		// query the database & scan the values into the variables, ** ALWAYS ** check for errors
		query := "SELECT id, username, password, created_at FROM users WHERE id = ?"

		// db.QueryRow() queries a specific row in the database
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
				log.Fatal(err)
		}

		fmt.Println(id, username, password, createdAt)
}

// Query all users
{
		// create variables to store data from the DB ; here a user
		type user struct {
				id        int
				username  string
				password  string
				createdAt time.Time
		}

		// select all rows / the rows you want to query; check for an error
		rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)

		if err != nil {
				log.Fatal(err)
		}
		// Close the rows after query & the function finishes executing
		defer rows.Close()

		// use a slice to store all the users
		var users []user
		// loop through each row from the query to load up user info
		for rows.Next() {
				var u user // user variable

				// Scan the row for user data & load into variables, check for error
				err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
				if err != nil {
						log.Fatal(err)
				}
				// add this user row to the slice of users
				users = append(users, u)
		}
		// check if there is an error while querying rows(empty row e.t.c)
		if err := rows.Err(); err != nil {
				log.Fatal(err)
		}

		fmt.Printf("%#v", users)
}

// Deleting a user from the database
{
	// query the db, check for error
		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
		if err != nil {
				log.Fatal(err)
		}
}
}