package gojsonenv

import (
	"encoding/json"

	"os"
)

var DefaultEnvFile = ".env.json"

// LoadEnv reads .env.json file and map its content to the env interface struct
func LoadEnv(env interface{}) {
	rawContent, _ := os.ReadFile(DefaultEnvFile)
	json.Unmarshal(rawContent, &env)
}
