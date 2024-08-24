package db

import (
	"bot/pkg/config"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (
	dbDriver        = "mysql"
	defaultNickname = "Anonymous"

	dbLoggerSlowSQLTreshold   = time.Second * 3
	dbLoggerLevel             = logger.Warn
	dbLoggerColorEnabled      = true
	dbLoggerIgnoreNotFoundErr = true
)

type dbHandler struct {
	conn *sql.DB
	gorm *gorm.DB

	tablesPrefix string
}

type DB interface{}

func New(cfg config.DBConfig, tablePrefix string) (DB, error) {
	lg := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             dbLoggerSlowSQLTreshold,
			LogLevel:                  dbLoggerLevel,
			IgnoreRecordNotFoundError: dbLoggerIgnoreNotFoundErr,
			Colorful:                  dbLoggerColorEnabled,
		},
	)

	var err error
	var conn *sql.DB
	var connErr error
	if conn, err = sql.Open(dbDriver, GetDBConnectionURI(cfg)); err != nil {
		return nil, fmt.Errorf("open sqldb connection: %v", err)
	}
	if connErr != nil {
		return nil, fmt.Errorf("db conn error: %w", err)
	}

	//conn.SetMaxOpenConns(cfg.MaxOpenConns)

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("ping db: %w", err)
	}

	mysqlConnConfig := mysql.New(mysql.Config{
		Conn: conn,
	})
	prefix := tablePrefix + "_"
	gormConfig := &gorm.Config{
		SkipDefaultTransaction:   true,
		DisableNestedTransaction: true,
		Logger:                   lg,
		NowFunc: func() time.Time {
			ti, err := time.LoadLocation(cfg.Location)
			if err != nil {
				panic(err)
			}

			return time.Now().In(ti)
		},
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: prefix,
		},
	}

	gormConn, err := gorm.Open(mysqlConnConfig, gormConfig)
	if err != nil {
		return nil, fmt.Errorf("open gorm conn: %w", err)
	}

	// migrate
	for _, prefab := range models {
		if err := gormConn.AutoMigrate(prefab); err != nil {
			return nil, fmt.Errorf("failed to migrate: %w", err)
		}
	}

	return &dbHandler{
		conn:         conn,
		gorm:         gormConn,
		tablesPrefix: prefix,
	}, nil
}

func GetDBConnectionURI(cfg config.DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?timeout=%dms&parseTime=true",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.ConnTimeoutMS,
	)
}
