package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
	"time"
)

type (
	Config struct {
		Environment string
		Postgres    PostgresConfig
		HTTP        HTTPConfig
		GRPC        GRPCConfig
		GRPCRA      GRPCConfigRA
		GRPCCS      GRPCConfigCS
		GRPCPS      GRPCConfigPS
	}

	PostgresConfig struct {
		Port     string
		Sslmode  string
		Host     string
		Username string ``
		Dbname   string
		Password string
	}

	HTTPConfig struct {
		Host               string
		Port               string
		ReadTimeout        time.Duration
		WriteTimeout       time.Duration
		MaxHeaderMegabytes int
	}

	GRPCConfig struct {
		Port string
	}

	GRPCConfigRA struct {
		Host string
		Port string
	}

	GRPCConfigCS struct {
		Host string
		Port string
	}

	GRPCConfigPS struct {
		Host string
		Port string
	}
)

func Init(configsDir string) (*Config, error) {

	if err := parseConfigFile(configsDir); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func parseConfigFile(folder string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.MergeInConfig()
}

func unmarshal(cfg *Config) error {

	if err := viper.UnmarshalKey("db", &cfg.Postgres); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("grpc", &cfg.GRPC); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("grpcRA", &cfg.GRPCRA); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("grpcCS", &cfg.GRPCCS); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("grpcPS", &cfg.GRPCPS); err != nil {
		return err
	}

	return nil
}

func setFromEnv(cfg *Config) {
	_ = godotenv.Load()
	cfg.Postgres.Password = os.Getenv("DB_PASSWORD")

}
