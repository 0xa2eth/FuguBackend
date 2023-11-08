package config

import (
	"FuguBackend/pkg/snowflake"
	"io/ioutil"
	"log"

	"github.com/naoina/toml"
	"github.com/naoina/toml/ast"
	"github.com/spf13/viper"
)

// CommonConfig Common
type CommonConfig struct {
	Version   string
	IsDebug   bool
	LogLevel  string
	LogPath   string
	StartTime string
}

// ServerConf config struct
type ServerConf struct {
	Addr string
}

type RedisConf struct {
	Addr         string `toml:"addr"`
	Pass         string `toml:"pass"`
	Db           int    `toml:"db"`
	MaxRetries   int    `toml:"maxretries"`
	PoolSize     int    `toml:"poolsize"`
	MinIdleConns int    `toml:"minidleconns"`
}

// MySQLConf mysql配置
type MySQLConf struct {
	Read struct {
		User     string `toml:"user" json:"user"` // 用户
		Password string `toml:"pass" json:"pass"` // 密码
		Host     string `toml:"host" json:"host"` // 地址
		Port     string `toml:"port" json:"port"` // 端口
		Database string `toml:"database" json:"database"`
	} `toml:"read" json:"read"`
	Write struct {
		User     string `toml:"user" json:"user"` // 用户
		Password string `toml:"pass" json:"pass"` // 密码
		Host     string `toml:"host" json:"host"` // 地址
		Port     string `toml:"port" json:"port"` // 端口
		Database string `toml:"database" json:"database"`
	} `toml:"write" json:"write"`
	Base struct {
		MaxIdleConns    int    `toml:"maxidleconn" mapstructure:"maxidleconn" json:"maxidleconn"`             // 最大空闲连接数
		MaxOpenConns    int    `toml:"maxopenconn" mapstructure:"maxopenconn" json:"maxopenconn"`             // 最大打开连接数
		ConnMaxLifeTime int    `toml:"connmaxlifetime" mapstructure:"connmaxlifetime" json:"connmaxlifetime"` // 连接复用时间
		LogLevel        string `toml:"loglevel" mapstructure:"loglevel" json:"loglevel"`                      // 日志级别，枚举（info、warn、error和silent）

	} `toml:"base" json:"base"`
}

type EtcdConf struct {
	Host1 string `toml:"host1" json:"host1"` //节点1
	Host2 string `toml:"host2" json:"host2"` //节点2
	Host3 string `toml:"host3" json:"host3"` //节点3
}

type AwsConf struct {
	Bucket    string `toml:"bucket" json:"bucket"`
	AccessKey string `toml:"access_key" json:"access_key"`
	SecretKey string `toml:"secret_key" json:"secret_key"`
	Region    string `toml:"region" json:"region"`     //
	Endpoint  string `toml:"endpoint" json:"endpoint"` //
}

type JwtConf struct {
	Issuer         string `toml:"issuer" json:"issuer"`
	SecretKey      string `toml:"secret_key" json:"secret_key"`
	ExpirationTime int64  `toml:"expiration_time" json:"expiration_time"`
}

type TwitterConf struct {
	AccessKey string `toml:"access_key" json:"access_key"`
}
type HashIds struct {
	Secret string `toml:"secret" json:"secret"`
	Length int    `toml:"length" json:"length"`
}
type Mail struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
	User string `toml:"user"`
	Pass string `toml:"pass"`
	To   string `toml:"to"`
}
type Language struct {
	Local string `toml:"local"`
}

// Config ...
type Config struct {
	Common   *CommonConfig `toml:"common"`
	Server   *ServerConf   `toml:"server"`
	MySQL    *MySQLConf    `toml:"mysql"`
	Redis    *RedisConf    `toml:"redis"`
	Etcd     *EtcdConf     `toml:"etcd"`
	Aws      *AwsConf      `toml:"aws"`
	Jwt      *JwtConf      `toml:"jwt"`
	Twitter  *TwitterConf  `toml:"twitter"`
	HashIds  *HashIds      `toml:"hashids"`
	Language *Language     `toml:"language"`
	Mail     *Mail         `toml:"mail"`
}

// Conf ...
var Conf = &Config{}

// LoadConfig ...
func LoadConfig() {
	// init the new config params
	initConf()

	contents, err := ioutil.ReadFile("./config/fugucnf.toml")
	if err != nil {
		log.Fatal("[FATAL] load fugucnf.toml: ", err)
	}
	tbl, err := toml.Parse(contents)
	if err != nil {
		log.Fatal("[FATAL] parse fugucnf.toml: ", err)
	}
	// parse common config
	parseCommon(tbl)
	// init log
	InitLogger()
	// parse server config
	parseServer(tbl)

	//parse mysql config
	parseMsq(tbl)

	//parse redis config
	parseReds(tbl)

	parseEtcd(tbl)

	parseAws(tbl)

	parseJwt(tbl)

	parseTwitter(tbl)

	parseHashids(tbl)

	// snowflake
	snowflake.SonyFlakeInit("2023-11-11", 8)
}

func initConf() {
	Conf = &Config{
		Common:   &CommonConfig{},
		Server:   &ServerConf{},
		MySQL:    &MySQLConf{},
		Redis:    &RedisConf{},
		Etcd:     &EtcdConf{},
		Aws:      &AwsConf{},
		Jwt:      &JwtConf{},
		Twitter:  &TwitterConf{},
		HashIds:  &HashIds{},
		Language: &Language{},
		Mail:     &Mail{},
	}
}

func parseCommon(tbl *ast.Table) {
	if val, ok := tbl.Fields["common"]; ok {
		subTbl, ok := val.(*ast.Table)
		if !ok {
			log.Fatalln("[FATAL] : ", subTbl)
		}

		err := toml.UnmarshalTable(subTbl, Conf.Common)
		if err != nil {
			log.Fatalln("[FATAL] parseCommon: ", err, subTbl)
		}
	}
}

func parseServer(tbl *ast.Table) {
	if val, ok := tbl.Fields["ser"]; ok {
		subTbl, ok := val.(*ast.Table)
		if !ok {
			log.Fatalln("[FATAL] : ", subTbl)
		}

		err := toml.UnmarshalTable(subTbl, Conf.Server)
		if err != nil {
			log.Fatalln("[FATAL] parseServer: ", err, subTbl)
		}
	}
}

func parseMsq(tbl *ast.Table) {
	if val, ok := tbl.Fields["mysql"]; ok {
		subTbl, ok := val.(*ast.Table)
		if !ok {
			log.Fatalln("[FATAL] : ", subTbl)
		}

		err := toml.UnmarshalTable(subTbl, Conf.MySQL)
		if err != nil {
			log.Fatalln("[FATAL] parseMySQL: ", err, subTbl)
		}
	}
}

func parseReds(tbl *ast.Table) {
	if val, ok := tbl.Fields["redis"]; ok {
		subTbl, ok := val.(*ast.Table)
		if !ok {
			log.Fatalln("[FATAL] : ", subTbl)
		}

		err := toml.UnmarshalTable(subTbl, Conf.Redis)
		if err != nil {
			log.Fatalln("[FATAL] parseReds: ", err, subTbl)
		}
	}
}

func parseEtcd(tbl *ast.Table) {
	if val, ok := tbl.Fields["etcd"]; ok {
		subTbl, ok := val.(*ast.Table)
		if !ok {
			log.Fatalln("[FATAL] : ", subTbl)
		}

		err := toml.UnmarshalTable(subTbl, Conf.Etcd)
		if err != nil {
			log.Fatalln("[FATAL] parseEtcd: ", err, subTbl)
		}
	}
}

func parseAws(tbl *ast.Table) {
	if val, ok := tbl.Fields["aws"]; ok {
		subTbl, ok := val.(*ast.Table)
		if !ok {
			log.Fatalln("[FATAL] : ", subTbl)
		}

		err := toml.UnmarshalTable(subTbl, Conf.Aws)
		if err != nil {
			log.Fatalln("[FATAL] parseAws: ", err, subTbl)
		}
	}
}

func parseJwt(tbl *ast.Table) {
	if val, ok := tbl.Fields["jwt"]; ok {
		subTbl, ok := val.(*ast.Table)
		if !ok {
			log.Fatalln("[FATAL] : ", subTbl)
		}

		err := toml.UnmarshalTable(subTbl, Conf.Jwt)
		if err != nil {
			log.Fatalln("[FATAL] parseAws: ", err, subTbl)
		}
	}
}

func parseTwitter(tbl *ast.Table) {
	if val, ok := tbl.Fields["twitter"]; ok {
		subTbl, ok := val.(*ast.Table)
		if !ok {
			log.Fatalln("[FATAL] : ", subTbl)
		}

		err := toml.UnmarshalTable(subTbl, Conf.Twitter)
		if err != nil {
			log.Fatalln("[FATAL] parseAws: ", err, subTbl)
		}
	}
}
func parseHashids(tbl *ast.Table) {
	if val, ok := tbl.Fields["hashids"]; ok {
		subTbl, ok := val.(*ast.Table)
		if !ok {
			log.Fatalln("[FATAL] : ", subTbl)
		}

		err := toml.UnmarshalTable(subTbl, Conf.HashIds)
		if err != nil {
			log.Fatalln("[FATAL] parseHashids: ", err, subTbl)
		}
	}
}
func UnmarshalConfig(configFilePath string) (*Config, error) {
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType("toml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	config, err := DefaultConfig()
	if err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}
	return config, nil
}

func DefaultConfig() (*Config, error) {
	return &Config{}, nil
}
func Get() Config {
	return *Conf
}
