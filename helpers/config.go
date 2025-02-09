package helpers

import (
	"strconv"

	"github.com/joho/godotenv"
)

var Env = map[string]string{}

func SetupConfig() {
	var err error
	Env, err = godotenv.Read(".env")
	if err != nil {
		Logger.Fatal("failed to read env file: ", err)
	}
}

func GetEnv(key string, val string) string {
	result := Env[key]
	if result == "" {
		result = val
	}
	return result
}

func GetEnvInt(key string, defaultValue int) int {
	valueStr := GetEnv(key, strconv.Itoa(defaultValue))
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}
