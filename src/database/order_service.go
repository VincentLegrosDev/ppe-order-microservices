package database 


import (
	"time"
	"database/sql"
	"ppe4peeps.com/services/models"
	"ppe4peeps.com/services/database/internal"
)


type orderDBService struct{
	conn *sql.DB
}

func (db orderDBService) InsertProcessedOrder(event models.OrderEvent) error { 
	 _, err := db.conn.Exec("insert into orders.processedEvents (id, event_name, processed_time) values ($1, $2, $3)",
			event.EventBase.EventId,
			event.Topic(), 
			time.Now(),
		)
	if err != nil {
		return err 
	}

	return nil
} 

func (db orderDBService) OrderAlreadyProcess(event models.OrderEvent) (bool, error)  {
	var id string

	err := db.conn.QueryRow("select id from orders.processedEvents where id=$1 and event_name=$2", event.EventBase.EventId, event.Topic()).Scan(&id)

	if err != nil {
		return false, err
	}

	return len(id) > 0, nil
}

func NewDatabaseConn() orderDBService {
	return orderDBService {
		conn: database.InitDatabase(),
	}
}
