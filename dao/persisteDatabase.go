package dao

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "10.3.81.150"
	port     = 5432
	user     = "valemobi"
	password = "v4l3m0b1"
	dbname   = "marketdata"
)

// DbConnection !
func DbConnection() *sql.DB {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	return db

}

// InsertTest !
func InsertTest(idUser string, pathFoto string) {
	db := DbConnection()

	// close database
	defer db.Close()

	sqlStatement := "INSERT INTO tb_test_investor_custom_preferences (id_investor, path_foto) VALUES ($1, $2);"
	_, erro := db.Exec(sqlStatement, idUser, pathFoto)
	CheckError(erro)

}

// GetTest !
func GetTest() {
	db := DbConnection()
	defer db.Close()

	var id int
	var name string

	sqlStatement := "SELECT * from tb_test_investor_custom_preferences;"
	rows, err := db.Query(sqlStatement)

	CheckError(err)

	defer rows.Close()

	linha := 0
	for rows.Next() {
		switch err := rows.Scan(&id, &name); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:

			fmt.Println("Linha: ", linha)
			fmt.Printf("ID_INVESTOR = %d\n", id)
			fmt.Printf("PATH_FOTO = %s\n", name)
			linha++
			fmt.Println("***---------------------------***")
		default:
			CheckError(err)
		}
	}

}

// CheckError !
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
