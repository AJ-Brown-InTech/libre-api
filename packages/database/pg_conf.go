//postgrees db connection
package database
import (
	"fmt"
	"github.com/AJ-Brown-InTech/libre-ra/config"
	"github.com/AJ-Brown-InTech/libre-ra/packages/utils"
	_ "github.com/jackc/pgx/stdlib" //"github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"time"
	"strconv"
)



//db connectors DO NOT TOUCH
const (
	maxOpenConns    = 100
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)



//create a new db connection ALWAYS CLOSE WHEN USING so no idle connections are hanging around
func NewPsqlDb(c *config.Config, log utils.Logger)(*sqlx.DB, error){
	
	var db *sqlx.DB
	var err error

	//db config
	p := c.Postgres.PostgresqlPort
	 port, err := strconv.ParseInt(p, 10, 0)

	 if err != nil {
		fmt.Println("Error during conversion")
		return nil,err
	}

	//host := c.Postgres.PostgresqlHost  "host=%s
	dbname := c.Postgres.PostgresqlDbname 
	user := c.Postgres.PostgresqlUser
	password := c.Postgres.PostgresqlPassword
	driver := c.Postgres.PgDriver

	dataSourceName := fmt.Sprintf( "port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable", 
		//host,
		port,
		user,
		password,
		dbname,
	)

	db, err = sqlx.Connect(driver, dataSourceName) 
	if err != nil{
		log.Errorf("Postgres database connection failed. [ERROR]:%v", err)
		return nil,err
	}

	defer db.Close()
	db.SetMaxOpenConns(maxOpenConns)//not sure how many but from what i read 100 is max before performance becomes an issue
	db.SetMaxIdleConns(maxIdleConns)//idle just added a few incase some connections are hung up
	db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)

	// if err = db.Ping(); err != nil{
	// 	log.Errorf("Postgres database connection issue.[ERROR]:%v", err)
	// 	return nil,err
	// }
	
	log.Infof("Postgres database connection SUCCESS.[INFO]:%v", db) 
	return db, nil
}
