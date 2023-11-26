package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	DebugMode bool     `mapstructure:"debug_mode"`
	Http      Http     `mapstructure:"http"`
	Mysql     Mysql    `mapstructure:"mysql"`
	Delivery  Delivery `mapstructure:"delivery"`
	Rabbitmq  Rabbitmq `mapstructure:"rabbitmq"`
}

type Http struct {
	Url    string `mapstructure:"url"`
	PodUrl string `mapstructure:"pod_url"`
	Host   string `mapstructure:"host"`
	Port   int    `mapstrucuture:"port"`
}

type Mysql struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"username"`
	Database string `mapstructure:"database"`
	Password string `mapstructure:"password"`
}

type Rabbitmq struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`

	QueueName    string `mapstructure:"queue_name"`
	RoutingKey   string `mapstructure:"routing_key"`
	ExchangeName string `mapstructure:"exchange_name"`
}

type Delivery struct {
	EstimateUrl string `mapstructure:"estimate_url"`
}

func New(prefix string) (*Config, error) {
	v := viper.New()
	v.SetEnvPrefix(prefix)
	v.AutomaticEnv()
	v.SetConfigName(prefix + "_config")
	v.AddConfigPath(".")
	v.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	bindEnv(v)
	_ = v.ReadInConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Info("Config file changed:"+e.Name, "")
	})

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

func bindEnv(v *viper.Viper) {
	_ = v.BindEnv("http.url", "http_url")
	_ = v.BindEnv("http.host", "http_host")
	_ = v.BindEnv("http.port", "http_port")

	_ = v.BindEnv("mysql.host", "mysql_host")
	_ = v.BindEnv("mysql.port", "mysql_port")
	_ = v.BindEnv("mysql.username", "mysql_username")
	_ = v.BindEnv("mysql.password", "mysql_password")
	_ = v.BindEnv("mysql.database", "mysql_database")

	_ = v.BindEnv("delivery.estimate_url", "delivery_estimate_url")

	_ = v.BindEnv("rabbitmq.username", "rabbitmq_username")
	_ = v.BindEnv("rabbitmq.password", "rabbitmq_password")
	_ = v.BindEnv("rabbitmq.host", "rabbitmq_host")
	_ = v.BindEnv("rabbitmq.port", "rabbitmq_port")
	_ = v.BindEnv("rabbitmq.queue_name", "rabbitmq_queue_name")
	_ = v.BindEnv("rabbitmq.routing_key", "rabbitmq_routing_key")
	_ = v.BindEnv("rabbitmq.exchange_name", "rabbitmq_exchange_name")
}
