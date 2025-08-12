package config

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func LoadConfig() Config {
	return Config{
		DBUser:     "root",
		DBPassword: "root",
		DBHost:     "127.0.0.1",
		DBPort:     "3306",
		DBName:     "testdb",
	}
}