package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// CF -> for use configs model
	CF = &Configs{}
)

// DatabaseConfig database config model
type DatabaseConfig struct {
	Host         string `mapstructure:"HOST"`
	Port         int    `mapstructure:"PORT"`
	Username     string `mapstructure:"USERNAME"`
	Password     string `mapstructure:"PASSWORD"`
	DatabaseName string `mapstructure:"DATABASE_NAME"`
	DriverName   string `mapstructure:"DRIVER_NAME"`
	Timeout      string `mapstructure:"TIMEOUT"`
}

// Configs config models
type Configs struct {
	App struct {
		ProjectID   string `mapstructure:"PROJECT_ID"`
		Host        string `mapstructure:"HOST"`
		WebBaseURL  string `mapstructure:"WEB_BASE_URL"`
		APIBaseURL  string `mapstructure:"API_BASE_URL"`
		Release     bool   `mapstructure:"RELEASE"`
		Port        string `mapstructure:"PORT"`
		Environment string `mapstructure:"ENVIRONMENT"`
	} `mapstructure:"APP"`
	Swagger struct {
		Title       string   `mapstructure:"TITLE"`
		Version     string   `mapstructure:"VERSION"`
		Host        string   `mapstructure:"HOST"`
		BaseURL     string   `mapstructure:"BASE_URL"`
		Description string   `mapstructure:"DESCRIPTION"`
		Schemes     []string `mapstructure:"SCHEMES"`
		Enable      bool     `mapstructure:"ENABLE"`
	} `mapstructure:"SWAGGER"`
	SMTP struct {
		Host        string `mapstructure:"HOST"`
		Port        int    `mapstructure:"PORT"`
		Username    string `mapstructure:"USERNAME"`
		Password    string `mapstructure:"PASSWORD"`
		Sender      string `mapstructure:"SENDER"`
		SenderAlias string `mapstructure:"SENDER_ALIAS"`
	} `mapstructure:"SMTP"`
	SQL DatabaseConfig `mapstructure:"SQL"`
}

// InitConfig init config
func InitConfig(configPath string) error {
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName("config")
	v.AutomaticEnv()
	v.SetConfigType("yml")

	if err := v.ReadInConfig(); err != nil {
		logrus.Error("read config file error:", err)
		return err
	}

	if err := bindingConfig(v, CF); err != nil {
		logrus.Error("binding config error:", err)
		return err
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		if err := bindingConfig(v, CF); err != nil {
			logrus.Error("binding error:", err)
			return
		}
	})

	return nil
}

// bindingConfig binding config
func bindingConfig(vp *viper.Viper, cf *Configs) error {
	if err := vp.Unmarshal(&cf); err != nil {
		logrus.Error("unmarshal config error:", err)
		return err
	}

	return nil
}
