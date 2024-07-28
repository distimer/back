package configs

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

var Env ServerEnv

type ServerEnv struct {
	LogLevel           string `env:"LOG_LEVEL"`
	DBHost             string `env:"DB_HOST"`
	DBPort             string `env:"DB_PORT"`
	DBUser             string `env:"DB_USER"`
	DBPass             string `env:"DB_PASS"`
	DBName             string `env:"DB_NAME"`
	JWTSecret          string `env:"JWT_SECRET"`
	JWTExpire          int    `env:"JWT_EXPIRE"`
	AppleClientID      string `env:"APPLE_CLIENT_ID"`
	GoogleClientID     string `env:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `env:"GOOGLE_CLIENT_SECRET"`
}

func getEnvStr(envName string) string {
	osEnv := os.Getenv(envName)
	if osEnv == "" {
		panic(fmt.Sprintf("Environment variable %s not set", envName))
	}
	return osEnv
}

func getEnvInt(envName string) int {
	osEnv := getEnvStr(envName)
	intEnv, err := strconv.Atoi(osEnv)
	if err != nil {
		panic(fmt.Sprintf("Environment variable %s is not an integer", envName))
	}
	return intEnv

}

func LoadEnv() {
	t := reflect.TypeOf(ServerEnv{})

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		envName := field.Tag.Get("env")
		switch field.Type.Kind() {
		case reflect.String:
			reflect.ValueOf(&Env).Elem().Field(i).SetString(getEnvStr(envName))
		case reflect.Int:
			reflect.ValueOf(&Env).Elem().Field(i).SetInt(int64(getEnvInt(envName)))
		default:
			panic(fmt.Sprintf("Unsupported type %s", field.Type.Kind()))
		}
	}
}
