package commands

import (
	"github.com/codegangsta/cli"
)

// Commands will return all supported commands
func Commands() []cli.Command {
	return []cli.Command{
		{
			Name:    "list-projects",
			Aliases: []string{"lp"},
			Usage:   "Lists all projects of a Gerrit instance",
			Action:  ListProjects,
		},
		{
			Name:    "crawl",
			Aliases: []string{"c"},
			Usage:   "Crawls a Gerrit instance",
			Action:  Crawl,
		},
	}
}
