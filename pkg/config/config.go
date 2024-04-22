package config

import "github.com/joho/godotenv"

func LoadAllConfigs(envFile string, mode string) {

	if mode != "production" {
		err := godotenv.Load(envFile)
		if err != nil {
			panic("Error loading.env file Err: " + err.Error())
		}
	}
	LoadAppConfig()
	LoadDBConfig()
}
