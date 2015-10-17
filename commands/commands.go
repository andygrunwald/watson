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
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:   "concurrent, c",
					Value:  200,
					Usage:  "Number of concurrent HTTP(S) calls",
					EnvVar: "WATSON_CONCURRENT",
				},
				cli.StringFlag{
					Name:   "storage, s",
					Value:  "",
					Usage:  "DSN of the storage backend",
					EnvVar: "WATSON_STORAGE",
				},
				cli.StringFlag{
					Name:   "identity-storage, is",
					Value:  "",
					Usage:  "DSN of the storage backend for identities",
					EnvVar: "WATSON_IDENTITY_STORAGE",
				},
			},
		},
	}
}
