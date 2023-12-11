package sqldb

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

func (db *Database) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		db.Host, db.User, db.Password, db.Dbname, db.Port)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func dbConnection() {
	db := &Database{
		Host:     "testtstsst",
		Port:     5432,
		User:     "postgres",
		Password: "password",
		Dbname:   "short_url_service",
	}

	gormDB, err := db.Connect()
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database connection established", gormDB)
	// Use gormDB for database operations
}
