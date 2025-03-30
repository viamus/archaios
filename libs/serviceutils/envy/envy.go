package envy

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

func LoadDotEnvIfDev() {
	once.Do(func() {
		if os.Getenv("ENV") != "production" {
			_ = godotenv.Load()
		}
	})
}

func Get(Key string) string {
	return os.Getenv(Key)
}
