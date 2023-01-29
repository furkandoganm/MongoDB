package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMongoURI() string {
	err := godotenv.Load() //.env dosyasını okur ve bu işlemi env'ye yükler

	if err != nil {
		log.Fatalln("error .env")
	}

	mongoURI := os.Getenv("MONGOURI")
	return mongoURI
}
