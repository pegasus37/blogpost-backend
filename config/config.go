package config

type Configuration struct {
	DB *DBConfig
}

type DBConfig struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
	Driver       string
}

func GetConfig() *Configuration {
	return &Configuration{
		DB: &DBConfig{
			Host:         "35.240.204.148",
			Port:         5432,
			Username:     "pegasus37",
			Password:     "PegasuSt33ga7",
			DatabaseName: "pegasus37",
			Driver:       "postgres",
		},
	}
}
