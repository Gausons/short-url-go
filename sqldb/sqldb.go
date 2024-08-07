package sqldb

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	gorm.Model
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

func (db *Database) connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		db.Host, db.User, db.Password, db.Dbname, db.Port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func DbConnection() {
	db := &Database{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		Dbname:   "short_url_service",
	}

	gormDB, err := db.connect()
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database connection established", gormDB)
	// Use gormDB for database operations
}

type ShortUrl struct {
	gorm.Model
	ShortUrl string `gorm:"type:varchar(8);"`
	LongUrl  string `gorm:"type:varchar(4096);"`
}

func FindShortUrl() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&ShortUrl{})

	// Create
	db.Create(&ShortUrl{ShortUrl: "srturl", LongUrl: "lngurl"})

	// Read
	var url ShortUrl
	db.First(&url, 1)                         // find url with integer primary key
	db.First(&url, "short_url = ?", "srturl") // find url with short_url srturl

	// Update - update url's long_url to lngurl2
	db.Model(&url).Update("LongUrl", "lngurl2")

	// Delete - delete url
	db.Delete(&url)
}
