package config

import (
	"fmt"
	"github.com/joho/godotenv"
)

func Init(envName string) {
	godotenv.Load(fmt.Sprintf(".env%s", envName))
}
