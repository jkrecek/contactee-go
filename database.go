package main

import (
	"github.com/jinzhu/gorm"
	"net/url"
)

var (
	instance *DB
)

type DB struct {
	*gorm.DB
}

func DbInstanceOptional(dbType, dbDsn string) (*DB, error) {
	if instance == nil {
		tempInstance, err := establishConnection(dbType, dbDsn)
		if err != nil {
			return nil, err
		}

		instance = &DB{DB: tempInstance}
		instance.Init(false)
	}

	return instance, nil
}

func establishConnection(dbType, dsn string) (db *gorm.DB, err error) {
	ud, err := url.Parse(dsn)
	if err != nil {
		return
	}

	q := ud.Query()
	q.Set("charset", "utf8")
	q.Set("parseTime", "True")
	q.Set("loc", "Local")
	ud.RawQuery = q.Encode()

	db, err = gorm.Open(dbType, ud.String())
	return
}

func (db *DB) Init(debug bool) {
	db.SingularTable(true)
	//db.InstantSet("gorm:save_associations", false)
	if debug {
		db.LogMode(true)
	}
}

func (db *DB) Migrate() {
	db.prepareDatabaseStructure()
}

var allDatabaseEntities = []interface{}{
	&DbContact{},
	&DbEmail{},
	&DbAddress{},
	&DbPhone{},
}

func (db *DB) IsDatabasePrepared() bool {
	return db.isStructureValid()
}

func (db *DB) prepareDatabaseStructure() {
	db.AutoMigrate(allDatabaseEntities...)
}

func (db *DB) isStructureValid() bool {
	for _, ent := range allDatabaseEntities {
		if !db.HasTable(ent) {
			return false
		}
	}

	return true
}
