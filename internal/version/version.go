package version

import (
        "context"
        "encoding/csv"
	"fmt"
	"log"
	"os"
	"os/exec"

        "github.com/google/go-github/v55/github"
)

func fetchLatestRelease(owner, repo string) (*github.RepositoryRelease, error) {
	client := github.NewClient(nil)
	ctx := context.Background()
	latest, _, err := client.Repositories.GetLatestRelease(ctx, owner, repo)
	return latest, err
}

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

		latest, err := fetchLatestRelease(item[1], item[2])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Local version of %s %slast version in GitHub repo is %s\nurl %s\n\n",
		    item[0],
		    string(stdout),
		    latest.GetTagName(),
		    latest.GetHTMLURL())

		}
}
