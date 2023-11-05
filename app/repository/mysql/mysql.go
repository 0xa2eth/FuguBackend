package mysql

import (
	"fmt"
	"time"

	"FuguBackend/config"
	"FuguBackend/pkg/errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Predicate is a string that acts as a condition in the where clause
type Predicate string

var (
	EqualPredicate              = Predicate("=")
	NotEqualPredicate           = Predicate("<>")
	GreaterThanPredicate        = Predicate(">")
	GreaterThanOrEqualPredicate = Predicate(">=")
	SmallerThanPredicate        = Predicate("<")
	SmallerThanOrEqualPredicate = Predicate("<=")
	LikePredicate               = Predicate("LIKE")
)

var _ Repo = (*dbRepo)(nil)

type Repo interface {
	i()
	GetDbR() *gorm.DB
	GetDbW() *gorm.DB
	DbRClose() error
	DbWClose() error
}

type dbRepo struct {
	DbR *gorm.DB
	DbW *gorm.DB
}

func New() (Repo, error) {
	cfg := config.Get().MySQL
	dbr, err := dbConnect(
		cfg.Read.User,
		cfg.Read.Password,
		fmt.Sprintf("%v:%v", cfg.Read.Host, cfg.Read.Port),
		cfg.Read.Database,
	)
	if err != nil {
		return nil, err
	}

	dbw, err := dbConnect(
		cfg.Write.User,
		cfg.Write.Password,
		fmt.Sprintf("%v:%v", cfg.Write.Host, cfg.Write.Port),
		cfg.Write.Database,
	)
	if err != nil {
		return nil, err
	}

	return &dbRepo{
		DbR: dbr,
		DbW: dbw,
	}, nil
}

func (d *dbRepo) i() {}

func (d *dbRepo) GetDbR() *gorm.DB {
	return d.DbR
}

func (d *dbRepo) GetDbW() *gorm.DB {
	return d.DbW
}

func (d *dbRepo) DbRClose() error {
	sqlDB, err := d.DbR.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (d *dbRepo) DbWClose() error {
	sqlDB, err := d.DbW.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// GetGormConfig 获取GORM相关配置
func getGormConfig() *gorm.Config {
	gc := &gorm.Config{
		QueryFields: true, // 根据字段名称查询
		PrepareStmt: true, // 缓存预编译语句
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 数据表名单数
		},
		NowFunc: func() time.Time {
			return time.Now() // 当前时间载入时区
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建外键约束
		//Logger: logger.Default.LogMode(logger.Info), // 日志配置
	}

	return gc
}

func dbConnect(user, pass, addr, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		user,
		pass,
		addr,
		dbName,
		true,
		"Local")
	// GORM MySQL相关配置
	my := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         255,  // string类型字段默认长度
		DisableDatetimePrecision:  true, // 禁用datetime精度
		DontSupportRenameIndex:    true, // 禁用重命名索引
		DontSupportRenameColumn:   true, // 禁用重命名列名
		SkipInitializeWithVersion: true, // 禁用根据当前mysql版本自动配置
	}

	db, err := gorm.Open(mysql.New(my), getGormConfig())

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("[db connection failed] Database name: %s", dbName))
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	cfg := config.Get().MySQL

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDB.SetMaxOpenConns(cfg.Base.MaxOpenConns)

	// 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	sqlDB.SetMaxIdleConns(cfg.Base.MaxIdleConns)

	// 设置最大连接超时
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(cfg.Base.ConnMaxLifeTime))

	// 使用插件
	db.Use(&TracePlugin{})

	return db, nil
}
