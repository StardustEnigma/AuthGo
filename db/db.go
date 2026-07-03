package db

import (
	"database/sql"
	"fmt"

	_"github.com/lib/pq"
)
var Db *sql.DB
func Connection()error{
	Connecstring := "host=localhost port=5432 user=postgress password= atharva@2004 dbName=Auth sslmode=disable"

	db,err := sql.Open("postgres",Connecstring);
	if(err != nil){
		return err
	}
	Db = db
	fmt.Println("Connected to db")
	return nil
}