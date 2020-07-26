package conn

import (
	"fmt"
	"log"
	"restapi/config"
	"sync"

	"github.com/jinzhu/gorm"
	//MSSQL
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

const (
	dbtype = "mssql"
)

//DB database structure
type DB struct {
	*gorm.DB
}

var dbInstance *DB

func connectDB(config *config.DBConfig) error {
	connString := fmt.Sprintf("server=%s; port=%d; database=%s;", config.Server, config.Port, config.DbName)
	conn, err := gorm.Open(dbtype, connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
		return err
	}
	fmt.Println("Database connection established")
	dbInstance = &DB{conn}
	return nil
}

//ConnectDB get database instance
func ConnectDB(config *config.DBConfig) *DB {

	var connDBOnce sync.Once
	connDBOnce.Do(func() {
		_ = connectDB(config)
	})
	return dbInstance
}

//AutoMigration :
func AutoMigration(db *DB) {
	db.AutoMigrate()
}
