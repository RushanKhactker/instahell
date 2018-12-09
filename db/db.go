package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const(
	DB_USER = "postgres"
	DB_PASSWORD = "postgr111es"
	DB_NAME = "xproj"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "xproj"
	password = ""
	dbname   = ""
)

type Log struct{
	Id int64			`json:"xId"`
	Date string 		`json:"xDate"`
	Msg string 			`json:"xMsg"`
}

func (x *Log) New () {

}

func Connect (){
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")


	sqlStatement := `insert into xLogs(                                                                                                                                                                                                                                                                             id, msg, dt) values ($1,$2,now())`
	_, err = db.Exec(sqlStatement, 1, "12345")
	if err != nil {
		panic(err)
	}

	//r, err := db.Query("select * from x.xlogs")
	//if err != nil {
	//	fmt.Println("Fuck Pizdec")
	//}
	//fmt.Println(r)
	//defer r.Close()


}