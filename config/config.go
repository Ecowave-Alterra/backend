package config

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func GetConfig() Configuration {

	configuration := Configuration{
		DB_USERNAME: "root",
		DB_PASSWORD: "",
		DB_PORT:     "3306",
		DB_HOST:     "localhost",
		DB_NAME:     "ecowave_db",
		// DB_USERNAME: os.Getenv("DB_USERNAME"),
		// DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		// DB_PORT:     os.Getenv("DB_PORT"),
		// DB_HOST:     os.Getenv("DB_HOST"),
		// DB_NAME:     os.Getenv("DB_NAME"),
	}

	return configuration
}
