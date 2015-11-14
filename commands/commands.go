package commands

import (
	"github.com/codegangsta/cli"
)

// Commands will return all supported commands
func Commands() []cli.Command {
	return []cli.Command{
		ListProjectsCommandDefinition(),
		CrawlCommandDefinition(),
	}
}
