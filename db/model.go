package db

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)



type XLog struct {
	Id 		uuid.UUID
	User 	string
	Date	time.Time
	Msg		string
}

const(
	SqlXlogInsert =
		`insert into xLog(xId,xUser,xDate,xMsg) 
		 values ($1,$2,now(),$3)`
	SqlXRequestInsert = ``)

func (x *XLog) Log_Insert (db *sql.DB) error {
	_, err := db.Exec(SqlXlogInsert, x.Id, x.User, x.Date, x.Msg)
	if err != nil {
		return err
	}
	return nil
}

//пока заглушка
type XRequest struct {
	xTime	time.Time
	xUser	string
	xMethod	string
	xData	string
}
func (x *XRequest) Request_Insert (db *sql.DB) error {
	return nil
}
