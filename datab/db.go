package datab

import (
  "fmt"
  "database/sql"
  _ "github.com/lib/pq"
)

// ToDo: This method required adding a trust record to pg_hba.conf
// Is there a better / more secure way?
const (
  host     = "127.0.0.1"
  port     = 5432
  user     = "dcrtest"
  // ToDo: moving password out of source?
  password = "testPassword"
  dbname   = "decred_test"
)

// Database struct for the db pool
type Database struct {
  pool *sql.DB
}

// Connect connects to the sql db
func (db *Database) Connect() { 
  
  var err error

  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
  //fmt.Println("Hello. Test output: ", psqlInfo)

  // validate the arguments to open connection to db
  db.pool, err = sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }

  // ToDo: Question - proper way to handle closing db at end of program?
  //defer db.pool.Close()

  // ping the connection info to actually open the db
  err = db.pool.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")
}

// CheckUserTable checks if table exists in the db
func (db *Database) CheckUserTable() error {
	sqlStatement := `SELECT id FROM users LIMIT 0;`
	_, err := db.pool.Query(sqlStatement)
	return err
}

// InsertRecord inserts a record and prints the id
func (db *Database) InsertRecord() {
	sqlStatement := `
		INSERT INTO users (age, email, first_name, last_name)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
	id := 0
	record := db.pool.QueryRow(sqlStatement, 30, "jon@calhoun.io", "Jonathan", "Calhoun")
	err := record.Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}

// UpdateRecord updates an existing record and prints the id
func (db *Database) UpdateRecord(id int) {
	sqlStatement := `
		UPDATE users
		SET first_name = $2, last_name = $3
		WHERE id = $1
		RETURNING id;`
	var returnedID int
	record := db.pool.QueryRow(sqlStatement, id, "NewFirst", "NewLast")
	err := record.Scan(&returnedID)
	if err != nil {
		panic(err)
	}
	fmt.Println(returnedID)
}

// DeleteRecord updates an existing record and prints the id
func (db *Database) DeleteRecord(id int) {

	sqlStatement := `
		DELETE FROM users
		WHERE id = $1
		RETURNING id;`
	var returnedID int
	record := db.pool.QueryRow(sqlStatement, id)
	err := record.Scan(&returnedID)
	if err != nil {
		panic(err)
	}
	fmt.Println(returnedID)
}

// ReadDB returns request id's info
func (db *Database) ReadDB(index int) string {
	type User struct {
		ID        int
		//Age       int
		//FirstName string
		//LastName  string
		Email     string
	}
	
	sqlStatement := `SELECT id, email FROM users WHERE id=$1;`
	var user User
	row := db.pool.QueryRow(sqlStatement, index)
	err := row.Scan(&user.ID, &user.Email)
	switch err {
		case sql.ErrNoRows:
			return fmt.Sprintf("No record with id %d found", index)
		case nil:
			return fmt.Sprintf("%d - %s", user.ID, user.Email)
		default:
			panic(err)
	}
}

// ManyRecords prints out multiple rows from the table
func (db *Database) ManyRecords() {
	rows, err := db.pool.Query("SELECT id, first_name FROM users LIMIT $1", 3)
  if err != nil {
    panic(err)
  }
  defer rows.Close()
  for rows.Next() {
    var id int
    var firstName string
    err = rows.Scan(&id, &firstName)
    if err != nil {
      panic(err)
    }
    fmt.Println(id, firstName)
  }
  // get any error encountered during iteration
  err = rows.Err()
  if err != nil {
    panic(err)
  }
}