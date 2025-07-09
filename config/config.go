package config

func NewConfig() Config {
	return Config{
		Server:   ServerConfig{},
		Logger:   LoggerConfig{},
		Postgres: PostgresConfig{},
		Telegram: TelegramConfig{},
	}
}

type Config struct {
	Server   ServerConfig
	Logger   LoggerConfig
	Postgres PostgresConfig
	Telegram TelegramConfig
}

type ServerConfig struct {
	Port string `env:"SERVER_PORT" envDefault:"8080"`
}

type LoggerConfig struct {
	Level    string `env:"LOG_LEVEL" envDefault:"info"`
	Encoding string `env:"LOG_ENCODING" envDefault:"json"`
}

type PostgresConfig struct {
	Database string `env:"POSTGRES_DATABASE" envDefault:"postgres"`
	Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	User     string `env:"POSTGRES_USER" envDefault:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" envDefault:"postgres"`
	Port     string `env:"POSTGRES_PORT" envDefault:"5432"`
	SSLMode  string `env:"POSTGRES_SSL_MODE" envDefault:"disable"`
}

type TelegramConfig struct {
	BotToken string `env:"TG_BOT_TOKEN" envDefault:"7613001634:AAHxu4xEbS2QMlImAZGv2aVjWTQTkkwfXn8"`
	MyID     string `env:"TG_MY_ID" envDefault:"1494574505"`
}
