package utils

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

func GetEnv(envStr string) (string, bool) {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Println(err)
		}
	})

	secret, ok := os.LookupEnv(envStr)
	return secret, ok
}
