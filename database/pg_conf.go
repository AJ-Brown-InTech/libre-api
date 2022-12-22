//postgrees db connection
package database
import (
	"fmt"
	"github.com/AJ-Brown-InTech/libre-api/config"
	_ "github.com/lib/pq"
	"database/sql"
	"time"
)

//db connectors DO NOT TOUCH
const (
	maxOpenConns    = 100
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

//create a new db connection ALWAYS CLOSE WHEN USING so no idle connections are hanging around
func NewPsqlDb(c *config.Config)(*sql.DB, error){
	
	host := c.Postgres.PostgresqlHost
	port := c.Postgres.PostgresqlPort
	dbname := c.Postgres.PostgresqlDbname
	user := c.Postgres.PostgresqlUser
	password := c.Postgres.PostgresqlPassword

	dataSourceName := fmt.Sprintf("host=%s, port=%s, dbname=%s, user=%s, password=%s, sslmode=disable", 
		host,
		port,
		dbname,
		user,
		password,
	)
	
	 db, err := sql.Open("postgres", dataSourceName )
	 if err != nil {
		return  nil, err
	}
	
	 db.SetMaxOpenConns(maxOpenConns)//not sure how many but from what i read 100 is max before performance becomes an issue
	 db.SetMaxIdleConns(maxIdleConns)//idle just added a few incase some connections are hung up
	 db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	 db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)
	
	err = db.Ping() //ping db server to verify connection
	if err != nil {
		return nil, err
	}
	return db, nil
}
