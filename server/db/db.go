package db

import (
	"fmt"
	"path/filepath"
	"runtime"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	fmt.Println("Connecting to database...")
	DB, err = gorm.Open(sqlite.Open(getDbPath()), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database: ", err)
	}

	fmt.Println("Auto migrate schemas...")
	autoMigrateAll()
}

// Otherwise database would be generated in main folder.
// This function ensures that database is generated in folder of the caller (db.go)
func getDbPath() string {
	_, currFile, _, _ := runtime.Caller(0)
	dbDir := filepath.Dir(currFile)
	dbPath := filepath.Join(dbDir, "database.db")

	return dbPath
}

func autoMigrateAll() {
	DB.AutoMigrate(&TemperatureReadout{})
}
