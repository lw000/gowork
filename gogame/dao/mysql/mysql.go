package moymysql

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"time"

	log "github.com/alecthomas/log4go"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"
)

type TyMysql struct {
	Db *sql.DB
}

var (
	ins *TyMysql
	mu  sync.Mutex
)

func SharedTymysqlInstance() *TyMysql {
	if ins == nil {
		mu.Lock()
		defer mu.Unlock()

		if ins == nil {
			ins = &TyMysql{}
		}
	}

	return ins
}

func (this *TyMysql) OpenSQL(cfg map[string]interface{}) (*TyMysql, error) {
	var (
		err error

		username     string
		password     string
		server       string
		database     string
		network      string = "tcp"
		port         int    = 3306
		MaxOpenConns int    = 10
		MaxOdleConns int    = 2
	)

	if v, ok := cfg["database"]; ok {
		database = v.(string)
	} else {
		return nil, errors.New("missing param database")
	}

	if v, ok := cfg["username"]; ok {
		username = v.(string)
	} else {
		return nil, errors.New("username is empty")
	}

	if v, ok := cfg["password"]; ok {
		password = v.(string)
	} else {
		return nil, errors.New("missing param password")
	}

	if v, ok := cfg["server"]; ok {
		server = v.(string)
	} else {
		return nil, errors.New("missing param server")
	}

	if v, ok := cfg["port"]; ok {
		port = v.(int)
	}

	if v, ok := cfg["network"]; ok {
		network = v.(string)
	}

	if v, ok := cfg["MaxOpenConns"]; ok {
		MaxOpenConns = v.(int)
	}

	if v, ok := cfg["MaxOdleConns"]; ok {
		MaxOdleConns = v.(int)
	}

	dns := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=true", username, password, network, server, port, database)

	this.Db, err = sql.Open("mysql", dns)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	this.Db.SetConnMaxLifetime(time.Hour * 5)
	this.Db.SetMaxOpenConns(MaxOpenConns)
	this.Db.SetMaxIdleConns(MaxOdleConns)

	if err := this.Db.Ping(); err != nil {
		log.Error(err)
		return nil, err
	}

	return this, nil
}

func (this *TyMysql) CloseSQL() {
	this.Db.Close()
}
