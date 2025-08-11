package config

type Config struct {
	DBConnString   string
}

func LoadConfig() Config {

	return Config {
		DBConnString: "DSN=MYSQL_ODBC;UID=root;PWD=root",
	}
}