package config

import "github.com/joho/godotenv"

func LoadAllConfigs(envFile string) {

	err := godotenv.Load(envFile)
	if err != nil {
		panic("Error loading.env file")
	}

	LoadAppConfig()
	LoadDBConfig()
}
