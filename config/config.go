package config

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() *Config {
	return &Config{
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUser:     "sonya",
		DBPassword: "postgres",
		DBName:     "movie_app",
	}
}
