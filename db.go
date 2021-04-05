package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Log struct {
	Data      string
	CreatedAt string
}

type DB struct {
	db *gorm.DB
}

func initDB(file string) *DB {
	db := DB{}

	db.db, _ = gorm.Open(sqlite.Open(file), &gorm.Config{})
	db.db.AutoMigrate(&Log{})

	return &db
}

func (db *DB) writeLog(data, time string) {
	newLog := Log{Data: data, CreatedAt: time}
	db.db.Create(&newLog)
}

func (db *DB) readLogs(box *[]Log) {
	db.db.Select("Data", "CreatedAt").Find(&box)
}
