package env

import "github.com/joho/godotenv"

func MustLoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
