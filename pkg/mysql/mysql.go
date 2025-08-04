package mysql

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true",
		getEnv("DB_USER", "root"),
		getEnv("DB_PASS", "yourpassword"),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_NAME", "xyz_finance"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek DB: ", err)
	}
	return db
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
