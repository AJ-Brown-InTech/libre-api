//postgrees db connection
package database
import (
	"fmt"
	"time"
	"github.com/AJ-Brown-InTech/libre-api/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/stdlib" // pgx driver"
)

const (
	maxOpenConns    = 100
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

func NewPsqlDb(c *config.Config)(*sqlx.DB, error){

	var dataSourceName string = fmt.Sprintf("host=%s, port=%s, dbname=%s, user=%s, password=%s, sslmode=disable", 
	c.Postgres.PostgresqlHost,
	string(c.Postgres.PostgresqlPort),
	c.Postgres.PostgresqlDbname,
	c.Postgres.PostgresqlUser,
	c.Postgres.PostgresqlPassword,
	)
	
	 db, err := sqlx.Open(c.Postgres.PgDriver, dataSourceName )
	 if err != nil {
		fmt.Println(nil)
		fmt.Println(err)
		return  nil, err
	}
	
	db.SetMaxOpenConns(maxOpenConns)//not sure how many but from what i read 100 is max before performance becomes an issue
	db.SetMaxIdleConns(maxIdleConns)//idle just added a few incase some connections are hung up
	db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)
	
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println(db)
	return db, nil
}
