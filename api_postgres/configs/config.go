package configs

import "github.com/spf13/viper"

//Variável criada apontando para config, para poder acessar
//os dados do arquivo
var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	viper.SetDefault("api.port", "8000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func Load() error {

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)
	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}
	println("Aqui")
	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
