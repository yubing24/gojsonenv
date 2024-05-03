# gojsonenv

Read a JSON-formatted .env file into a struct

# Installation

Add this to a project:
```sh
go get github.com/yubing24/gojsonenv
```

# Usage
1. Define your environment in a JSON file:
```json
{
   "DB_CONNECTION": "postgres://yourdbconnectionstring",
   "REDIS_CONNECTION": "redis://yourredisconnectionstring"
}
```

By default, save it to `.env.json` file in the same directory with your `main.go`.

2. In your application, define your application configuration struct:
```go
package main

import (
	"fmt"
	"github.com/yubing24/gojsonenv"
)

type AppConfig struct {
	DbConnection    string `json:"DB_CONNECTION"`
	RedisConnection string `json:"REDIS_CONNECTION"`
}

func main() {
	env := AppConfig{}
	gojsonenv.LoadEnv(&env)

	fmt.Printf("DbConnection: %v\n", env.DbConnection)
	fmt.Printf("RedisConnection: %v\n", env.RedisConnection)
}
```

If you prefer to use a different environment file or location, set the package variable `DefaultEnvFile`
to a different value, such as `.env.dev.json`.
