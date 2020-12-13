package databases

import (
	"github.com/go-redis/redis/v7"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"net"
	"time"
)

type RedisType = string

const (
	Simple   RedisType = "simple"
	Sentinel           = "sentinel"
	Cluster            = "cluster"
)

type RedisConfig struct {
	Host             string    `mapstructure:"host"`
	Port             string    `mapstructure:"port"`
	Db               int       `mapstructure:"db"`
	IdleSize         int       `mapstructure:"idlesize"`
	MaxActive        int       `mapstructure:"maxactive"`
	Password         string    `mapstructure:"password"`
	IdleTimeout      int       `mapstructure:"idletimeout"`
	ReadTimeout      int       `mapstructure:"readtimeout"`
	WriteTimeout     int       `mapstructure:"writetimeout"`
	ConnectTimeout   int       `mapstructure:"connecttimeout"`
	SentinelAddrs    []string  `mapstructure:"sentinel_addrs"`
	SentinelPassword string    `mapstructure:"sentinel_password"`
	MasterName       string    `mapstructure:"master_name"`
	Type             RedisType `mapstructure:"type"`
}

func NewRedisClient(viper *viper.Viper) (client *redis.Client, err error) {
	var config RedisConfig
	if err := viper.UnmarshalKey("redis", &config); err != nil {
		return nil, err
	}
	switch config.Type {
	case Sentinel:
		client = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:       config.MasterName,
			SentinelAddrs:    config.SentinelAddrs,
			SentinelPassword: config.SentinelPassword,
			Password:         config.Password,
			DB:               config.Db,
		})
	case Cluster:
		panic("not implement yet!")
	default:
		client = redis.NewClient(&redis.Options{
			Addr:         net.JoinHostPort(config.Host, config.Port),
			ReadTimeout:  time.Duration(config.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(config.WriteTimeout) * time.Second,
			DialTimeout:  time.Duration(config.ConnectTimeout) * time.Second,
			IdleTimeout:  time.Duration(config.IdleTimeout) * time.Second, // important!!! idle timeout must little .....
			DB:           config.Db,
			Password:     config.Password,
		})
	}
	return
}

var RedisProvideSet = wire.NewSet(NewRedisClient)
