package database 

import (
	"time"
	"database/sql"
	"ppe4peeps.com/services/models"
	"ppe4peeps.com/services/database/internal"
)

type orderDBService [T models.Event ] struct{
	conn *sql.DB
}

func (db orderDBService[T]) InsertProcessedEvent(event T) error { 
	 _, err := db.conn.Exec("insert into orders.processedEvents (id, event_name, processed_time) values ($1, $2, $3)",
			event.Id(),
			event.Topic(), 
			time.Now(),
		)

	if err != nil {
		return err 
	}

	return nil
} 

func (db orderDBService[T]) EventAlreadyProcess(event T) (bool, error)  {  
	var id string

	err := db.conn.QueryRow("select id from orders.processedEvents where id=$1 and event_name=$2", event.Id(), event.Topic()).Scan(&id)

	if err != nil && err != sql.ErrNoRows {  
		return false, err
	} 

	if err == sql.ErrNoRows {
		return false, nil 
	} else {
		return true, nil
	}

}

func NewDatabaseConn[T models.Event]() orderDBService[T] {
	return orderDBService[T] {
		conn: database.InitDatabase(),
	}
}
