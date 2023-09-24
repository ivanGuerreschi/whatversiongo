package version

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func GetVersion(csvFile string) {
	file, err := os.Open(csvFile)
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records")
	}

	for _, item := range records[1:] {
		cmd := exec.Command(item[0], "--version")
		stdout, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(string(stdout))
	}
}
