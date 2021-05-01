package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Log struct {
	Data      string
	CreatedAt string
}

type DB struct {
	db *gorm.DB
}

func initDB(dsm string) *DB {
	db := DB{}

	db.db, _ = gorm.Open(mysql.Open(dsm), &gorm.Config{})
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
