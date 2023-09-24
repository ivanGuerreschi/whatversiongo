package utility

import (
	"log"
	"os"
)

func GetHome() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return homeDir
}
