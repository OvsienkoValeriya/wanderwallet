package config

import (
	"flag"
	"os"
	"sync"
)

type Config struct {
	RunAddress  string
	DatabaseURI string
}

var (
	cfg  *Config
	once sync.Once
)

func Init() *Config {
	once.Do(func() {
		addrFlag := flag.String("a", "", "address and port")
		dbFlag := flag.String("d", "", "Postgres URI")
		flag.Parse()

		runAddr := "localhost:3000"
		dbURI := "postgres://postgres:postgres@localhost:5432/Wander_Wallet?sslmode=disable"

		if *addrFlag != "" {
			runAddr = *addrFlag
		} else if env := os.Getenv("RUN_ADDRESS"); env != "" {
			runAddr = env
		}
		if *dbFlag != "" {
			dbURI = *dbFlag
		} else if env := os.Getenv("DATABASE_URI"); env != "" {
			dbURI = env
		}
		cfg = &Config{RunAddress: runAddr, DatabaseURI: dbURI}
	})
	return cfg
}

func Get() *Config {
	if cfg == nil {
		panic("config not initialized")
	}
	return cfg
}
