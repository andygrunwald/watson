package commands

import (
	"fmt"
	"github.com/andygrunwald/go-gerrit"
	"github.com/codegangsta/cli"
	"os"
)

// Crawl crawls a Gerrit instance
func Crawl(c *cli.Context) {
	instance := c.GlobalString("instance")

	client, err := gerrit.NewClient(instance, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Crawl all the things ...", client)
	os.Exit(0)
}
