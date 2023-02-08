package orm

import (
	"fmt"
	"time"

	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type (
	mysqlConfig struct {
		address  string
		username string
		password string
		dbName   string

		maxIdleConnection int
		maxOpenConnection int
		connMaxLifetime   time.Duration

		logger           Logger
		logMode          bool
		logSlowThreshold time.Duration
	}

	MysqlConfig interface {
		Apply(c *mysqlConfig)
	}
)

func defaultMysqlConfig() (mysqlConfig, error) {
	logger, err := zapLog.NewZap()
	if err != nil {
		return mysqlConfig{}, err
	}

	return mysqlConfig{
		maxIdleConnection: 10,
		maxOpenConnection: 200,
		connMaxLifetime:   time.Hour,

		logger:           logger,
		logMode:          false,
		logSlowThreshold: 200 * time.Millisecond,
	}, nil
}

func (c mysqlConfig) generateURI() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		c.username, c.password, c.address, c.dbName)
}

func (c mysqlConfig) ToGormConfig() *gorm.Config {
	conf := gorm.Config{}
	if c.logMode {
		logConf := gormLogger.Config{
			Colorful: true,
			LogLevel: gormLogger.Info,
		}

		if c.logSlowThreshold != 0 {
			logConf.SlowThreshold = c.logSlowThreshold
		}

		conf.Logger = gormLogger.New(c.logger, logConf)
	}

	return &conf
}

type withAddress struct {
	address string
}

func (w withAddress) Apply(c *mysqlConfig) {
	c.address = w.address
}

func WithAddress(address string) MysqlConfig {
	return withAddress{address}
}

type withUsernamePassword struct {
	username string
	password string
}

func (w withUsernamePassword) Apply(c *mysqlConfig) {
	c.username = w.username
	c.password = w.password
}

func WithUsernamePassword(username, password string) MysqlConfig {
	return withUsernamePassword{username, password}
}

type withDatabaseName struct {
	dbName string
}

func (w withDatabaseName) Apply(c *mysqlConfig) {
	c.dbName = w.dbName
}

func WithDatabaseName(dbName string) MysqlConfig {
	return withDatabaseName{dbName}
}

type withMaxIdleConnection struct {
	maxIdleConnection int
}

func (w withMaxIdleConnection) Apply(c *mysqlConfig) {
	c.maxIdleConnection = w.maxIdleConnection
}

func WithMaxIdleConnection(maxIdleConnection int) MysqlConfig {
	return withMaxIdleConnection{maxIdleConnection}
}

type withMaxOpenConnection struct {
	maxOpenConnection int
}

func (w withMaxOpenConnection) Apply(c *mysqlConfig) {
	c.maxOpenConnection = w.maxOpenConnection
}

func WithMaxOpenConnection(maxOpenConnection int) MysqlConfig {
	return withMaxOpenConnection{maxOpenConnection}
}

type withConnMaxLifetime struct {
	connMaxLifetime time.Duration
}

func (w withConnMaxLifetime) Apply(c *mysqlConfig) {
	c.connMaxLifetime = w.connMaxLifetime
}

func WithConnMaxLifetime(connMaxLifetime time.Duration) MysqlConfig {
	return withConnMaxLifetime{connMaxLifetime}
}

type withLogMode struct {
	logMode bool
}

func (w withLogMode) Apply(c *mysqlConfig) {
	c.logMode = w.logMode
}

func WithLogMode(logMode bool) MysqlConfig {
	return withLogMode{logMode}
}

type withLogger struct {
	logger Logger
}

func (w withLogger) Apply(c *mysqlConfig) {
	c.logger = w.logger
	c.logMode = true
}

func WithLogger(logger Logger) MysqlConfig {
	return withLogger{logger}
}

type withLog struct {
	logger           Logger
	logSlowThreshold time.Duration
}

func (w withLog) Apply(c *mysqlConfig) {
	c.logger = w.logger
	c.logSlowThreshold = w.logSlowThreshold
	c.logMode = true
}

func WithLog(logger Logger, slowThreshold time.Duration) MysqlConfig {
	return withLog{logger, slowThreshold}
}
