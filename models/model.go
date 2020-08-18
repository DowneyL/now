package models

import (
	"fmt"
	"github.com/DowneyL/now/packages/configs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var (
	db  *gorm.DB
	wdb *gorm.DB
	err error
)

type BaseModel struct {
	ID          uint      `gorm:"primary_key;type:bigint(20) unsigned not null auto_increment;comment:'主键ID'" json:"id"`
	CreatedTime time.Time `gorm:"not null;default:current_timestamp;comment:'创建时间'" json:"created_time"`
	UpdatedTime time.Time `gorm:"not null;default:current_timestamp on update current_timestamp;comment:'更新时间'" json:"updated_time"`
}

type Model struct {
	BaseModel
	DeletedTime *time.Time `sql:"index;not null;default:'1970-01-01 00:00:00';comment:'删除时间'" json:"-"`
}

func SetUp() {
	config := configs.New()
	readDatabase := config.ReadDatabase
	readConnect := getDBConnectInfo(readDatabase)
	db, err = gorm.Open(readDatabase.Type, readConnect)
	if err != nil {
		log.Fatalf("read db connect error:%v", err)
	}
	setup(db, config)

	writeDatabase := config.WriteDatabase
	writeConnect := getDBConnectInfo(writeDatabase)
	wdb, err = gorm.Open(writeDatabase.Type, writeConnect)
	if err != nil {
		log.Fatalf("write db connect error:%v", err)
	}
	setup(wdb, config)
}

func Close(db *gorm.DB) {
	defer db.Close()
}

func setup(db *gorm.DB, config *configs.Config) {
	db.SingularTable(true)
	if config.Server.Mode != gin.ReleaseMode {
		db.LogMode(true)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func getDBConnectInfo(db configs.Database) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db.User, db.Password, db.Host, db.Port, db.Name)
}
