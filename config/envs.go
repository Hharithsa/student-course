package config

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	APIKey     string
}

var Envs = initializeConfig()

func initializeConfig() Config {
	return Config{
		PublicHost: "http://127.0.0.1",
		Port:       "5000",
		DBUser:     "root",
		DBPassword: "password",
		DBAddress:  "http://127.0.0.1:3306",
		DBName:     "localdb",
		APIKey:     "dummy-key",
	}
}
